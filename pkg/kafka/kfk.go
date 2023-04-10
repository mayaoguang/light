package kafka

import (
	"context"
	"fmt"
	kfk "github.com/segmentio/kafka-go"
	"time"
)

type Config struct {
	Addr []string
}

var cfg *Config

func Init(c *Config) {
	cfg = c
}

// NewWriter 异步写入, 高性能, WriteMessages不会阻塞, 错误异步处理.
// async true 异步写入, 高性能, 错误异步处理. false 同步写入, 性能较低, 能实时返回错误
func NewWriter(topic string, async bool, Completion func([]kfk.Message, error)) *kfk.Writer {
	w := &kfk.Writer{
		Addr:         kfk.TCP(cfg.Addr...),
		Topic:        topic,
		Balancer:     &kfk.LeastBytes{},
		Async:        async,
		Completion:   Completion,
		RequiredAcks: kfk.RequireOne,
	}

	return w
}

// Consumer 消息消费.
func Consumer(ctx context.Context, topic, groupId string, handle func(kfk.Message) error) error {
	r := kfk.NewReader(kfk.ReaderConfig{
		Brokers: cfg.Addr,
		Topic:   topic,
		MaxWait: 100 * time.Millisecond,
		GroupID: groupId,
	})
	defer r.Close()
	for {
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("fetch msg err: %+v", err)
		}
		if err = handle(msg); err != nil {
			return err
		}
		if err = r.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("commit msg err: %+v", err)
		}
	}
}
