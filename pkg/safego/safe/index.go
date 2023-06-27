package safe

import (
	"context"
	"fmt"
	mqSdk "github.com/aliyunmq/mq-http-go-sdk"
	"log"
	"runtime/debug"
	"sync"
)

// Go 安全go程.
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(fmt.Sprintf("recover err: %v", err))
				debug.PrintStack()
			}
		}()
		f()
	}()
}

func GoWithCtx(ctx context.Context, f func(ctx context.Context)) {
	go func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover err:%+v", err)
				debug.PrintStack()
			}
		}()
		f(ctx)
	}(ctx)
}

// GoWithField 安全go程且携带参数
func GoWithField(f func(val interface{}), val interface{}) {
	go func(val interface{}) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(fmt.Sprintf("recover err: %v", err))
				debug.PrintStack()
			}
		}()
		f(val)
	}(val)
}

func GoWithCtxWg(ctx context.Context, wg *sync.WaitGroup,
	f func(ctx context.Context, wg *sync.WaitGroup)) {
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(fmt.Sprintf("GoWithCtxWg recover err:%+v", err))
				debug.PrintStack()
			}
		}()
		f(ctx, wg)
	}(ctx, wg)
}

func GoWithMqMsg(ch <-chan struct{}, wg *sync.WaitGroup, message mqSdk.ConsumeMessageEntry,
	f func(ch <-chan struct{}, wg *sync.WaitGroup, message mqSdk.ConsumeMessageEntry)) {
	go func(ch <-chan struct{}, wg *sync.WaitGroup, message mqSdk.ConsumeMessageEntry) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(fmt.Sprintf("GoWithMqMsg recover err:%+v", err))
				debug.PrintStack()
			}
		}()
		f(ch, wg, message)
	}(ch, wg, message)
}
