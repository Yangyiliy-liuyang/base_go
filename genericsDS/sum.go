package genericsDS

type Number interface {
	int | uint | int64
}

func Sum[T Number](vals []T) T {
	var res T
	for _, v := range vals {
		res = res + v
	}
	return res
}
func Max[T Number](vals []T) T {
	t := vals[0]
	for i := 0; i < len(vals); i++ {
		if t < vals[i] {
			t = vals[i]
		}
	}
	return t
}
func Min[T Number](vals []T) T {
	t := vals[0]
	for i := 0; i < len(vals); i++ {
		if t > vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Find[T any](vals []T, filter func(t T) bool) T {
	for _, v := range vals {
		if filter(v) {
			return v
		}
	}
	var t T
	return t
}

func Insert[T any](idx int, val T, vals []T) []T {
	if idx < 0 || idx > len(vals) {
		panic("idx不合法")
	}
	//扩容
	vals = append(vals, val)
	for i := len(vals) - 1; i > idx; i-- {
		if i-1 > 0 {
			vals[i] = vals[i-1]
		}
	}
	vals[idx] = val
	return vals
}
