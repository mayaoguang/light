package generic

// MapEqual 比对两个map
func MapEqual[K comparable, V comparable](s, t map[K]V) (b bool) {
	if len(s) != len(t) {
		return false
	}
	for k, vs := range s {
		if vt, ok := t[k]; !ok || vs != vt {
			return false
		}
	}
	return true
}

// MapEqualPred 按照给定的比较方法比较两个map
func MapEqualPred[K comparable, V any](s, t map[K]V, equalFunc func(s, t V) bool) (b bool) {
	if len(s) != len(t) {
		return false
	}
	for k, vs := range s {
		if vt, ok := t[k]; !ok || !equalFunc(vs, vt) {
			return false
		}
	}
	return true
}

// MapKeys 获取keys
func MapKeys[K comparable, V any](m map[K]V) (r []K) {
	r = make([]K, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}
	return r
}

// MapValues 获取values切片
func MapValues[K comparable, V any](m map[K]V) (r []V) {
	r = make([]V, len(m))
	i := 0
	for _, v := range m {
		r[i] = v
		i++
	}
	return r
}

// MapValuesMap 获取values map
func MapValuesMap[K comparable, V comparable](m map[K]V) (r map[V]struct{}) {
	r = make(map[V]struct{}, len(m))
	for _, v := range m {
		r[v] = struct{}{}
	}
	return r
}

// MapContainKeys 查看一批key 是否在map 中
func MapContainKeys[K comparable, V any](m map[K]V, keys ...K) (b bool) {
	for _, k := range keys {
		if _, exists := m[k]; !exists {
			return false
		}
	}
	return true
}

// MapContainValues 查看一批values 是否在map中
func MapContainValues[K comparable, V comparable](m map[K]V, values ...V) (b bool) {
	vMap := MapValuesMap(m)
	for _, v := range values {
		if _, ok := vMap[v]; !ok {
			return false
		}
	}
	return true
}

// MapMerge 把t 合并到 s 中 并修改s 中key 相同的值，并保持t 不变
func MapMerge[K comparable, V any](s, t map[K]V) map[K]V {
	if t == nil {
		return s
	}
	if s == nil {
		s = make(map[K]V, len(t))
	}
	for k, v := range t {
		s[k] = v
	}
	return s
}

// MapGet 获取值
func MapGet[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		return val
	}
	return defaultVal
}

// MapPop 删除map中指定的key的值，并返回删除的值,如果不存在,返回默认值
func MapPop[K comparable, V any](m map[K]V, k K, defaultVal V) V {
	if val, ok := m[k]; ok {
		delete(m, k)
		return val
	}
	return defaultVal
}

// MapUnionKeys 一批map 取所有不重复的key
func MapUnionKeys[K comparable, V any](m map[K]V, ms ...map[K]V) (r []K) {
	r = MapKeys(m)
	for _, v := range ms {
		r = SliceUnion(r, MapKeys(v))
	}
	return r
}

func MapIntersectionKeys[K comparable, V any](m map[K]V, ms ...map[K]V) (r []K) {
	r = MapKeys(m)
	for _, v := range ms {
		r = SliceIntersect(r, MapKeys(v))
	}
	return
}
