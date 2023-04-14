package utils

import (
	"math/rand"
	"reflect"
	"time"
)

// InSlice val 是否在slice 中 类型和值都必须一样
// 只遍历1次，比转换成map再查快
func InSlice(val interface{}, slice interface{}) (exists bool, index int) {
	exists = false
	index = -1

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) == false {
			continue
		}
		index = i
		exists = true
		return
	}
	return
}

// SliceUnique 数组唯一
func SliceUnique(slice []string) (r []string) {
	m := make(map[string]struct{})
	for _, v := range slice {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		r = append(r, v)
	}
	return
}

// SliceShuffle 数组乱序
func SliceShuffle(slice interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	var ran = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := s.Len() - 1; i > 0; i-- {
		k := ran.Intn(i + 1)
		tmp := s.Index(k).Interface()
		s.Index(k).Set(s.Index(i))
		s.Index(i).Set(reflect.ValueOf(tmp))
	}
	return
}

// SliceStrIntersect 两个数组(字符串)取交集
func SliceStrIntersect(f, s []string) (r []string) {
	r = make([]string, 0)
	tmpMap := make(map[string]struct{})
	for _, v := range f {
		tmpMap[v] = struct{}{}
	}
	for _, v := range s {
		if _, ok := tmpMap[v]; ok {
			r = append(r, v)
		}
	}
	return
}

// SliceIntersect 两个数组(不限类型)取交集
func SliceIntersect(f, s interface{}) (r []interface{}) {
	if reflect.TypeOf(f).Kind() != reflect.Slice || reflect.TypeOf(s).Kind() != reflect.Slice {
		return
	}
	r = make([]interface{}, 0)
	tmpMap := make(map[interface{}]struct{})

	vf := reflect.ValueOf(f)
	for i := 0; i < vf.Len(); i++ {
		tmpMap[vf.Index(i).Interface()] = struct{}{}
	}
	vs := reflect.ValueOf(s)
	for i := 0; i < vs.Len(); i++ {
		if _, ok := tmpMap[vs.Index(i).Interface()]; ok {
			r = append(r, vs.Index(i).Interface())
		}
	}
	return
}

// SliceDiff 交集取反 (即在 f中且不在s的值+在s且不在f中的值)
func SliceDiff(f, s interface{}) (r []interface{}) {
	if reflect.TypeOf(f).Kind() != reflect.Slice || reflect.TypeOf(s).Kind() != reflect.Slice {
		return
	}
	r = make([]interface{}, 0)
	fMap, sMap := make(map[interface{}]struct{}), make(map[interface{}]struct{})
	vf := reflect.ValueOf(f)
	for i := 0; i < vf.Len(); i++ {
		fMap[vf.Index(i).Interface()] = struct{}{}
	}
	vs := reflect.ValueOf(s)
	for i := 0; i < vs.Len(); i++ {
		sMap[vs.Index(i).Interface()] = struct{}{}
	}

	for i := 0; i < vf.Len(); i++ {
		if _, ok := sMap[vf.Index(i).Interface()]; ok {
			continue
		}
		r = append(r, vf.Index(i).Interface())
	}

	for i := 0; i < vs.Len(); i++ {
		if _, ok := fMap[vs.Index(i).Interface()]; ok {
			continue
		}
		r = append(r, vs.Index(i).Interface())
	}
	return
}

// SortMerge 有序数组合并并去重
func SortMerge(s1, s2 []int) []int {
	l := len(s1) + len(s2)
	r := make([]int, 0, l)
	i, j := 0, 0
	for {
		if i == len(s1) {
			r = append(r, s2[j:]...)
			break
		}
		if j == len(s2) {
			r = append(r, s1[i:]...)
		}
		if s1[i] < s2[j] {
			r = append(r, s1[i])
			i++
		} else if s1[i] == s2[j] {
			j++
			continue
		} else {
			r = append(r, s2[j])
			j++
		}

	}
	return r
}
