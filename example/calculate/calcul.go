package main

import (
	"fmt"
	"light/pkg/calculate"
	"log"
)

// 一个简易的表达式计算器
func main() {
	express := []byte("5*6+3*4-1-3*4+10")
	result, err := calculate.Calculate(express)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("result:", result)
}
