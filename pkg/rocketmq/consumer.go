package rocketmq

import (
	"context"
	"fmt"
	mqSdk "github.com/aliyunmq/mq-http-go-sdk"
	"light/pkg/safego/safe"
	"log"
	"strings"
	"sync"
	"time"
)

type (
	Consumer struct {
		ConsumerCommon                                   // consumer 的配置
		C              mqSdk.MQConsumer                  // 消费者实体
		messageChan    chan mqSdk.ConsumeMessageResponse // 消息管道
		errChan        chan error                        // 错误管道
		processChan    chan mqSdk.ConsumeMessageEntry    // 处理消息的管道
		ackChan        chan string                       // ack 数据处理
		quit           chan struct{}                     // 处理消息的协程退出后ack线程才能退出
		ackArray       *safe.SafeSlice                   // 待ack数据的列表
	}
)

func NewConsumer(manger *ConsumerManger) *Consumer {
	c := &Consumer{
		ConsumerCommon: manger.ConsumerCommon,
	}
	return c
}

func (s *Consumer) Run(ctx context.Context, wg *sync.WaitGroup) {
	s.messageChan = make(chan mqSdk.ConsumeMessageResponse, s.chanNum)
	s.errChan = make(chan error, s.chanNum)
	s.processChan = make(chan mqSdk.ConsumeMessageEntry)
	s.ackChan = make(chan string, s.workerNum)
	s.quit = make(chan struct{})
	s.ackArray = safe.NewSafeSlice(s.ackNum)

	wg.Add(4)
	safe.GoWithCtxWg(ctx, wg, s.watch)
	safe.GoWithCtxWg(ctx, wg, s.handle)
	safe.GoWithCtxWg(ctx, wg, s.msgProcess)
	safe.GoWithCtxWg(ctx, wg, s.ackProcess)
}

// watch 监听消息
func (s *Consumer) watch(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			s.C.ConsumeMessage(s.messageChan, s.errChan, s.numOfMessages, s.waitSeconds)
		}
	}
}

// handle 处理
func (s *Consumer) handle(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case item := <-s.messageChan:
			for i := 0; i < len(item.Messages); i++ {
				s.processChan <- item.Messages[i]
			}
		case err := <-s.errChan:
			if !strings.Contains(err.Error(), "Message not exist") {
				log.Printf(fmt.Sprintf("call handle() errChan fail [%+v]", err))
			}
		case <-ctx.Done():
			// TODO 已经消费的数据没有回复ACK MQ会再次发送，这样退出不影响
			return
		}
	}
}

// msgProcess 单次消息多线程控制
func (s *Consumer) msgProcess(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	workerChan := make(chan struct{}, s.workerNum) // 控制任务并发的chan
	defer close(workerChan)
	workerWg := sync.WaitGroup{}
	for {
		select {
		case msg := <-s.processChan:
			workerWg.Add(1)
			workerChan <- struct{}{}
			safe.GoWithMqMsg(workerChan, &workerWg, msg, s.process)
		case <-ctx.Done():
			// 等当前消费的到的数据全部处理完
			workerWg.Wait()
			s.quit <- struct{}{}
			return
		}
	}
}

// process 处理消息
func (s *Consumer) process(ch <-chan struct{}, wg *sync.WaitGroup, message mqSdk.ConsumeMessageEntry) {
	defer func() {
		wg.Done()
		<-ch
	}()

	// 查询针对 tag 的处理方法
	handler, exists := s.handler[message.MessageTag]
	if !exists {
		_ = s.C.AckMessage([]string{message.ReceiptHandle})
		log.Println(fmt.Sprintf("未找到 tag 能够处理的方法 tag:%+v", message.MessageTag))
		return
	}
	if handler([]byte(message.MessageBody)) {
		s.ackArray.Append(message.ReceiptHandle)
		s.ackChan <- message.ReceiptHandle
	}
}

// ackProcess 数据消费完成 ack
func (s *Consumer) ackProcess(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	after := time.After(time.Duration(s.ackSecond) * time.Second)
	for {
		select {
		case <-s.ackChan:
			if s.ackArray.Len() >= s.ackNum {
				if err = s.ackArray.ProcessAndClear(s.C.AckMessage, DefaultAckNum); err != nil {
					log.Println(fmt.Sprintf("ack err:%+v,len(ackArray) = %d", err, s.ackArray.Len()))
					continue
				}
				after = time.After(time.Duration(s.ackSecond) * time.Second)
			}
		case <-after:
			if err = s.ackArray.ProcessAndClear(s.C.AckMessage, DefaultAckNum); err != nil {
				log.Println(fmt.Sprint("time ack err:%+v,len(ackArray) = %d", err, s.ackArray.Len()))
				continue
			}
			after = time.After(time.Duration(s.ackSecond) * time.Second)
		case <-ctx.Done():
			<-s.quit // 等消费的协程退出才能ack
			if err = s.ackArray.ProcessAndClear(s.C.AckMessage, DefaultAckNum); err != nil {
				log.Println(fmt.Sprintf("done ack err:%+v,ackArray = %+v", err, s.ackArray))
			}
			return
		}
	}
}
