package generic

import (
	"math/rand"
	"time"
)

// SliceEqual 比较两个slice 是否相等
func SliceEqual[T comparable](f, s []T) (b bool) {
	if len(f) != len(s) {
		return false
	}
	for i := range f {
		if f[i] != s[i] {
			return false
		}
	}
	return true
}

// SliceToMap 数组转map
func SliceToMap[T comparable](s []T) (r map[T]struct{}) {
	r = make(map[T]struct{}, len(s))
	for i := range s {
		r[s[i]] = struct{}{}
	}
	return
}

// InSlice val 是否在slice中
func InSlice[T comparable](val T, s []T) (b bool) {
	for i := range s {
		if val == s[i] {
			return true
		}
	}
	return false
}

// SliceToSet 数组唯一
func SliceToSet[T comparable](s []T) (r []T) {
	l := len(s)
	if l <= 1 {
		return s
	}
	m := make(map[T]struct{}, l)
	r = make([]T, 0, l)
	for i := range s {
		v := s[i]
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

	tmp := make(map[T]struct{}, lenF)
	r = make([]T, 0, Min(lenF, lenS))

	for i := range f {
		tmp[f[i]] = struct{}{}
	}
	for i := range s {
		v := s[i]
		if _, ok := tmp[v]; ok {
			r = append(r, v)
		}
	}
	return
}

// SliceUnion 两个数组(泛型)取并集
func SliceUnion[T comparable](f, s []T) (r []T) {
	lenF, lenS := len(f), len(s)
	if lenF == 0 {
		return SliceToSet(s)
	}
	if lenS == 0 {
		return SliceToSet(f)
	}

	tmp := make(map[T]struct{}, lenF+lenS)
	r = make([]T, 0, lenF+lenS)

	for i := range f {
		v := f[i]
		if _, ok := tmp[v]; ok {
			continue
		}
		tmp[v] = struct{}{}
		r = append(r, v)
	}
	for i := range s {
		v := s[i]
		if _, ok := tmp[v]; ok {
			continue
		}
		tmp[v] = struct{}{}
		r = append(r, v)
	}
	return
}

// SliceUnIntersect 交集取反 (即在 f中且s不在的值+在s且不在f中的值)
func SliceUnIntersect[T comparable](f, s []T) (r []T) {
	r = make([]T, len(f)+len(s))
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

// SliceReverse 翻转slice
func SliceReverse[T any](s []T) []T {
	if len(s) == 0 {
		return s
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
