package generic

// MapEqual 比对两个map
func MapEqual[K comparable, V comparable](s, t map[K]V) bool {
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
func MapEqualPred[K comparable, V any](s, t map[K]V, equalFunc func(s, t V) bool) bool {
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
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// MapValues 获取values
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// MapContainKeys 查看一批key 是否在map 中
func MapContainKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, k := range keys {
		if _, exists := m[k]; !exists {
			return false
		}
	}
	return true
}

// MapContainValues 查看一批values 是否在map中
func MapContainValues[K comparable, V comparable](m map[K]V, values ...V) bool {
	for _, v := range values {
		found := false
		for _, x := range m {
			if x == v {
				found = true
				break
			}
		}
		if !found {
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

// MapUnionKeys 获取一批map的key
func MapUnionKeys[K comparable, V any](m map[K]V, ms ...map[K]V) []K {
	keys := MapKeys(m)
	for _, v := range ms {
		keys = SliceUnion(keys, MapKeys(v))
	}
	return keys
}
