package main

import (
	"fmt"
	"light/pkg/logging"
	"sync"
)

var (
	wg = new(sync.WaitGroup)
)

func main() {
	defer func() {
		wg.Wait()
		logging.Sync()
	}()
	fmt.Println("main")
}
