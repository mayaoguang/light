package dp

import "fmt"

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

// Permutation 动态规划 无重复字符串的排列组合
func Permutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}
	list := []string{S}
	if len(S) == 1 {
		return list
	}
	l := len(S)
	for i := 0; i < l-1; i++ {
		for t := range list {
			for j := i + 1; j < l; j++ {
				s := []byte(list[t])
				s[i], s[j] = s[j], s[i]
				list = append(list, string(s))
			}
		}
	}
	return list
}

//permutation 递归 空间时间都比动态规划慢
func permutation(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}
	ret := []string{}
	for i := range S {
		tmp := S[:i] + S[i+1:]
		res := Permutation(tmp)
		for j := range res {
			ret = append(ret, fmt.Sprintf("%c%s", S[i], res[j]))
		}
	}
	return ret
}
