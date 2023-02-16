package generic

func Min[T NumberEx | ~string](v T, s ...T) (r T) {
	r = v
	for i := range s {
		if s[i] < r {
			r = s[i]
		}
	}
	return
}
