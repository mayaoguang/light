package generic

import (
	"strings"
	"unsafe"
)

// Filter filters slice elements with condition.
func Filter[T any](s []T, filterFunc func(t T) bool) []T {
	result := make([]T, 0, len(s))
	for i := range s {
		if filterFunc(s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

func FilterLt[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t < v })
}

func FilterLte[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t <= v })
}

func FilterGt[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t > v })
}

func FilterGte[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t >= v })
}

func FilterNe[T comparable](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t != v })
}

func FilterIn[T comparable](s []T, v ...T) []T {
	m := make(map[T]struct{}, len(v))
	for i := range v {
		m[v[i]] = struct{}{}
	}
	return Filter(s, func(t T) bool {
		_, ok := m[t]
		return ok
	})
}

func FilterNin[T comparable](s []T, v ...T) []T {
	m := make(map[T]struct{}, len(v))
	for i := range v {
		m[v[i]] = struct{}{}
	}
	return Filter(s, func(t T) bool {
		_, ok := m[t]
		return !ok
	})
}

// FilterLike 区分大小写
func FilterLike[T StringEx](s []T, v string) []T {
	if len(v) == 0 {
		return []T{}
	}
	return Filter(s, func(t T) bool {
		return strings.Contains(*(*string)(unsafe.Pointer(&t)), v)
	})
}

// FilterILike 不区分大小写
func FilterILike[T StringEx](s []T, v string) []T {
	if len(v) == 0 {
		return []T{}
	}
	v = strings.ToLower(v)
	return Filter(s, func(t T) bool {
		return strings.Contains(strings.ToLower(*(*string)(unsafe.Pointer(&t))), v)
	})
}
