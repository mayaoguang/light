package generic

import (
	"math/rand"
	"time"
)

// go v1.18 以上才能用
// 把之前slice 的方法用泛型重新实现了一下，基本上速度比使用interface快出10倍左右

// SliceToMap 数组转map
func SliceToMap[T comparable](s []T) (r map[T]struct{}) {
	r = make(map[T]struct{})
	for _, v := range s {
		r[v] = struct{}{}
	}
	return
}

// InSlice val 是否在slice中
func InSlice[T comparable](val T, s []T) (bool, int) {
	for i, v := range s {
		if val == v {
			return true, i
		}
	}
	return false, -1
}

// SliceUnique 数组唯一
func SliceUnique[T comparable](s []T) (r []T) {
	l := len(s)
	if l <= 1 {
		return s
	}
	m := make(map[T]struct{}, l)
	r = make([]T, 0, l)
	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		r = append(r, v)
	}
	return
}

// SliceShuffle 数组乱序
func SliceShuffle[T any](slice []T) {
	var ran = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		k := ran.Intn(i + 1)
		slice[i], slice[k] = slice[k], slice[i]
	}
	return
}

// SliceIntersect 两个数组(泛型)取交集
func SliceIntersect[T comparable](f, s []T) (r []T) {
	lenF, lenS := len(f), len(s)
	if lenF+lenS == 0 {
		return []T{}
	}

	r = make([]T, 0, Min(lenF, lenS))
	fMap := SliceToMap(f)
	for _, v := range s {
		if _, ok := fMap[v]; ok {
			r = append(r, v)
		}
	}
	return
}

// SliceDiff 交集取反 (即在 f中且不在s的值+在s且不在f中的值)
func SliceDiff[T comparable](f, s []T) (r []T) {
	r = make([]T, 0)
	fMap, sMap := SliceToMap(f), SliceToMap(s)
	for k := range fMap {
		if _, ok := sMap[k]; ok {
			continue
		}
		r = append(r, k)
	}
	for k := range sMap {
		if _, ok := fMap[k]; ok {
			continue
		}
		r = append(r, k)
	}
	return
}

// SliceUnion 两个数组(泛型)取并集
func SliceUnion[T comparable](f, s []T) (r []T) {
	lenF, lenS := len(f), len(s)
	if lenF == 0 {
		return SliceUnique(s)
	}
	if lenS == 0 {
		return SliceUnique(f)
	}

	tmp := make(map[T]struct{}, lenF+lenS)
	r = make([]T, 0, lenF+lenS)

	for _, v := range f {
		if _, ok := tmp[v]; ok {
			continue
		}
		tmp[v] = struct{}{}
		r = append(r, v)
	}
	for _, v := range s {
		if _, ok := tmp[v]; ok {
			continue
		}
		tmp[v] = struct{}{}
		r = append(r, v)
	}
	return
}
