package rocketmq

import (
	"encoding/json"
	"time"

	mqSdk "github.com/aliyunmq/mq-http-go-sdk"
)

type RocketMQ struct {
	cfg    *RocketMQConfig
	client mqSdk.MQClient
}

type RocketMQMsg struct {
	Topic       string `json:"topic,omitempty"`
	Message     string `json:"message,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Key         string `json:"key,omitempty"`
	IsDelayMsg  bool   `json:"is_delay_msg,omitempty"`
	DeliverTime int64  `json:"deliver_time,omitempty"`
	ErrMsg      string `json:"err_msg,omitempty"`
}

type Handle func(*RocketMQMsg)

func (rmq *RocketMQMsg) String() string {
	b, _ := json.Marshal(rmq)
	return string(b)
}

// Send 发送普通消息
func (s *RocketMQ) Send(topic, message, tag, key string, handle Handle) {
	producer := s.client.GetProducer(s.cfg.InstanceID, topic)
	go func() {
		for i := 1; i <= 10; i++ {
			_, err := producer.PublishMessage(mqSdk.PublishMessageRequest{
				MessageBody: message,             // 消息内容。
				MessageTag:  tag,                 // 消息标签。
				Properties:  map[string]string{}, // 消息属性。
				MessageKey:  key,
			})
			if err == nil {
				break
			}
			if err != nil && i == 10 && handle != nil {
				s.callBackDealMQMsg(topic, message, tag, key, err.Error(), 0, false, handle)
				break
			}
			time.Sleep(time.Second * 10)
		}
	}()
}

// NewRocketMQ 实例化一个
func NewRocketMQ(cfg RocketMQConfig) *RocketMQ {
	mq := &RocketMQ{cfg: &cfg,
		client: mqSdk.NewAliyunMQClient(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey, ""),
	}
	return mq
}

// TODO 不能大于40天
func (s *RocketMQ) SendDelayTimerMsg(topic, message, tag, key string,
	deliverTime int64, handle Handle) {
	producer := s.client.GetProducer(s.cfg.InstanceID, topic)
	go func() {
		for i := 1; i <= 10; i++ {
			_, err := producer.PublishMessage(mqSdk.PublishMessageRequest{
				MessageBody:      message,             // 消息内容。
				MessageTag:       tag,                 // 消息标签。
				Properties:       map[string]string{}, // 消息属性。
				MessageKey:       key,
				StartDeliverTime: deliverTime * 1000, // 毫秒级别单位
			})
			if err == nil {
				break
			} else if err != nil && i == 10 && handle != nil {
				s.callBackDealMQMsg(topic, message, tag, key, err.Error(), deliverTime, true, handle)
				break
			}
			time.Sleep(time.Second * 10)
		}
	}()
}
func (s *RocketMQ) callBackDealMQMsg(topic, message, tag, key, err string,
	deliverTime int64, isDelayMsg bool, handle Handle) {
	var info RocketMQMsg
	info.Topic = topic
	info.Message = message
	info.Tag = tag
	info.Key = key
	if isDelayMsg {
		info.DeliverTime = deliverTime
		info.IsDelayMsg = true
	}
	info.ErrMsg = err
	handle(&info)
}
