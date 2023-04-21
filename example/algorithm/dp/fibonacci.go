package dp

// 斐波那契数列

func Fibonacci(mapData map[int]int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if _, ok := mapData[n]; ok {
		return mapData[n]
	}
	mapData[n] = Fibonacci(mapData, n-1) + Fibonacci(mapData, n-2)
	return mapData[n]
}
