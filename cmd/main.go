package main

import (
	"context"
	"fmt"
	"light/pkg/safego/safe"
	"math/rand"
	"os"
	"runtime"
	"time"
)

// 多线程处理模型, 10个协程 处理数据
func main() {
	s := Su{
		messageChan: make(chan []int, 10),
		processChan: make(chan int),
		ackChan:     make(chan int),
	}
	ctx, _ := context.WithCancel(context.Background())
	safe.GoWithCtx(ctx, s.watch)
	safe.GoWithCtx(ctx, s.handle)
	safe.GoWithCtx(ctx, s.MesProcess)
	safe.GoWithCtx(ctx, s.AckProcess)
	for {
		time.Sleep(10 * time.Second)
	}
	os.Exit(1)
}

type Su struct {
	messageChan chan []int // 消息管道
	processChan chan int   // 处理消息的管道
	ackChan     chan int   // ack 数据处理
}

func (s *Su) watch(ctx context.Context) {
	i := 0
	for {
		fmt.Println("===========watch=====", i)
		tmp := i * 10
		s.messageChan <- []int{tmp + 0, tmp + 1, tmp + 2, tmp + 3, tmp + 4, tmp + 5, tmp + 6, tmp + 7, tmp + 8, tmp + 9}
		i++
		if i > 13 {
			return
		}
	}
}

func (s *Su) handle(ctx context.Context) {
	for {
		select {
		case item := <-s.messageChan:
			for i := 0; i < len(item); i++ {
				s.processChan <- item[i]
			}
		case <-ctx.Done():
			// TODO 已经消费的数据没有回复ACK MQ会再次发送，这样退出不影响
			return
		}
	}
}

func (s *Su) MesProcess(ctx context.Context) {
	c := make(chan struct{}, 100) // 控制任务并发的chan
	defer close(c)
	for {
		select {
		case msg := <-s.processChan:
			c <- struct{}{}
			go s.Process(c, msg)
		case <-ctx.Done():
			// TODO 已经消费的数据没有回复ACK MQ会再次发送，这样退出不影响
			return
		}
	}
}

func (s *Su) Process(ch <-chan struct{}, data int) {
	defer func() {
		<-ch
	}()
	fmt.Println(data, runtime.NumGoroutine())
	ran := rand.Int() % 100
	fmt.Println("ran", ran)
	time.Sleep(time.Duration(ran) * time.Millisecond)
	s.ackChan <- data
}

func (s *Su) AckProcess(ctx context.Context) {
	ackArray := make([]int, 0, 100)
	after := time.After(2 * time.Second)
	for {
		select {
		case ack := <-s.ackChan:
			ackArray = append(ackArray, ack)
			if len(ackArray) >= 100 {
				fmt.Println(len(ackArray), "ack", ackArray)
				ackArray = ackArray[0:0]
				after = time.After(2 * time.Second)
			}
		case <-after:
			fmt.Println(time.Now().Unix(), len(ackArray), "time ack", ackArray)
			ackArray = ackArray[0:0]
			after = time.After(2 * time.Second)
		case <-ctx.Done():
			return
		}
	}
}
