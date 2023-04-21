package algorithm

import (
	"fmt"
	"light/example/algorithm/dp"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tmp := make(map[int]int)
	fmt.Println(dp.Fibonacci(tmp, 10))
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(dp.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
func TestCommonSub(t *testing.T) {
	fmt.Println(dp.LongestCommonSubsequence("abc", "cdeabddc"))
}
