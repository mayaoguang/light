package kafka

import (
	"context"
	"fmt"
	kfk "github.com/segmentio/kafka-go"
	"testing"
	"time"
)

func conf() *Config {
	return &Config{Addr: []string{"127.0.0.1:9092"}}
}

var (
	topic = "topic"
)

func TestKfkWriter(t *testing.T) {
	Init(conf())
	w := NewWriter(topic, true, func(messages []kfk.Message, err error) {
		fmt.Println("async err", err, len(messages))
	})
	msg := make([]kfk.Message, 0)
	for i := 0; i < 10; i++ {
		msg = append(msg, kfk.Message{Value: []byte(fmt.Sprintf("num: %d", i))})
	}
	if err := w.WriteMessages(context.Background(), msg...); err != nil {
		t.Error(err.Error())
	}
	time.Sleep(1 * time.Second)
}

func TestKfkConsumer(t *testing.T) {
	Init(conf())
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))

	if err := Consumer(ctx, topic, "test", func(message kfk.Message) error {
		fmt.Println(string(message.Value))
		return nil
	}); err != nil {
		t.Error(err.Error())
	}
}
