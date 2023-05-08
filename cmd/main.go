package main

import (
	"fmt"
	"sync"
)

// 多线程处理模型
func main() {
	data := make(chan int, 20)
	results := make(chan int, 20)
	var wg sync.WaitGroup

	// 生产数据
	for i := 0; i < 20; i++ {
		data <- i
	}
	close(data)

	// 处理数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for d := range data {
				results <- process(d)
			}
		}()
	}

	// 等待所有任务完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 输出结果
	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("结果：", sum)
}

func process(d int) int {
	// 模拟处理数据
	fmt.Println(d)
	return d * d
}
