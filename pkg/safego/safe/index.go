package safe

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
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
