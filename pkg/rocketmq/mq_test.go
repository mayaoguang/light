package rocketmq

import (
	"context"
	"testing"
)

var TagFunc = map[string]func(data []byte) bool{
	"Check": func(data []byte) bool { return true },
}

func TestMqConsumer(t *testing.T) {
	rCfg := RocketMQConfig{
		Endpoint:   "http://xxx.com",
		AccessKey:  "xxx",
		SecretKey:  "xxxx",
		InstanceID: "MQ_INST_xxxxxxx",
	}
	cManger := NewConsumerManger(rCfg)
	cManger.AddTopicAndHandler("topic", TagFunc)
	cManger.SetGroup("group")
	cManger.Run(context.Background())
}
