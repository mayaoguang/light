package rocketmq

import (
	"context"
	"fmt"
	mqSdk "github.com/aliyunmq/mq-http-go-sdk"
	"log"
	"sync"
)

type (
	ConsumerManger struct {
		ConsumerCommon // worker 数量和message数量 配置
		cfg            *RocketMQConfig
		client         mqSdk.MQClient
		topics         []string
		handler        map[string]TopicHandler
	}
)

// NewConsumerManger 实例化一个consumer
func NewConsumerManger(cfg RocketMQConfig) *ConsumerManger {
	c := &ConsumerManger{
		ConsumerCommon: NewConsumerConf(),
		cfg:            &cfg,
		client:         mqSdk.NewAliyunMQClient(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey, ""),
		topics:         make([]string, 0),
		handler:        make(map[string]TopicHandler, 0),
	}

	return c
}

func (s *ConsumerManger) AddTopicAndHandler(topic string, h TopicHandler) *ConsumerManger {
	s.topics = append(s.topics, topic)
	s.handler[topic] = h
	return s
}

func (s *ConsumerManger) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for _, topic := range s.topics {
		consumer := NewConsumer(s)
		if _, ok := s.handler[topic]; !ok {
			log.Println(fmt.Sprintf("topic not has handler.topic:%+v", topic))
			continue
		}
		consumer.SetTopicAndHandler(topic, s.handler[topic]).SetGroup(s.group).SetTag(s.tag)
		consumer.C = s.client.GetConsumer(s.cfg.InstanceID, topic, s.group, s.tag)
		consumer.Run(ctx, &wg)
	}
	wg.Wait()
}
