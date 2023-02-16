package generic

func Min[T NumberEx | ~string](v T, s ...T) (min T) {
	min = v
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return
}

func MinInSlice[T NumberEx | ~string](s []T) (min T, err error) {
	if len(s) == 0 {
		return min, SliceEmpty
	}
	min = s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return
}

// MinInPred 自定义比较方法取最小值
func MinInPred[T any](s []T, lessFunc func(a, b T) bool) (min T, err error) {
	if len(s) == 0 {
		return min, SliceEmpty
	}
	min = s[0]
	for i := range s {
		if lessFunc(s[i], min) {
			min = s[i]
		}
	}
	return min, nil
}

func Max[T NumberEx | ~string](v T, s ...T) (max T) {
	max = v
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return
}

func MasInSlice[T NumberEx | ~string](s []T) (max T, err error) {
	if len(s) == 0 {
		return max, SliceEmpty
	}
	max = s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return
}

// MaxInPred 自定义比较方法取最大值
func MaxInPred[T any](s []T, lessFunc func(a, b T) bool) (max T, err error) {
	if len(s) == 0 {
		return max, SliceEmpty
	}
	max = s[0]
	for i := range s {
		if lessFunc(max, s[i]) {
			max = s[i]
		}
	}
	return max, nil
}
