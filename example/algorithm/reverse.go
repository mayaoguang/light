package algorithm

import (
	"strings"
)

// ReversePrefix 反转单词前缀
func ReversePrefix(word string, ch byte) string {
	index := strings.Index(word, string(ch))
	if index < 0 {
		return word
	}
	tmp := []byte(word)
	for i := 0; i <= index/2; i++ {
		tmp[i], tmp[index-i] = tmp[index-i], tmp[i]
	}
	return string(tmp)
}

// MaximumGap 164 最大间距
func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	tmpMap := make(map[int]struct{})
	s := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := tmpMap[nums[i]]; ok {
			continue
		}
		s = append(s, nums[i])
		tmpMap[nums[i]] = struct{}{}
	}
	QuickSort(s, 0, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		tmp := s[i] - s[i-1]
		if max < tmp {
			max = tmp
		}
	}
	return max
}

func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := arr[left]
		j := left
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		QuickSort(arr, left, j)
		QuickSort(arr, j+1, right)
	}
}

// SelectSort 选择排序
func SelectSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		min := arr[j]
		minIndex := j
		for i := j + 1; i < len(arr); i++ {
			if min > arr[i] {
				min = arr[i]
				minIndex = i
			}
		}
		if minIndex != j {
			arr[j], arr[minIndex] = arr[minIndex], arr[j]
		}
	}
}
