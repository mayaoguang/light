package utils

import (
	"math/rand"
	"time"
)

// go v1.18 以上才能用
// 把之前slice 的方法用泛型重新实现了一下，基本上速度比使用interface快出10倍左右

// GenericSliceToMap 数组转map
func GenericSliceToMap[T comparable](s []T) (r map[T]struct{}) {
	r = make(map[T]struct{})
	for _, v := range s {
		r[v] = struct{}{}
	}
	return
}

// GenericInSlice val 是否在slice中
func GenericInSlice[T comparable](val T, s []T) (bool, int) {
	for i, v := range s {
		if val == v {
			return true, i
		}
	}
	return false, -1
}

// GenericSliceUnique 数组唯一
func GenericSliceUnique[T comparable](s []T) (r []T) {
	m := make(map[T]struct{})
	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		r = append(r, v)
	}
	return
}

// GenericSliceShuffle 数组乱序
func GenericSliceShuffle[T any](slice []T) {
	var ran = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		k := ran.Intn(i + 1)
		slice[i], slice[k] = slice[k], slice[i]
	}
	return
}

// GenericSliceIntersect 两个数组(泛型)取交集
func GenericSliceIntersect[T comparable](f, s []T) (r []T) {
	r = make([]T, 0)
	fMap := GenericSliceToMap(f)
	for _, v := range s {
		if _, ok := fMap[v]; ok {
			r = append(r, v)
		}
	}
	return
}

// GenericSliceDiff 交集取反 (即在 f中且不在s的值+在s且不在f中的值)
func GenericSliceDiff[T comparable](f, s []T) (r []T) {
	r = make([]T, 0)
	fMap, sMap := GenericSliceToMap(f), GenericSliceToMap(s)
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
