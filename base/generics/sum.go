package generics

import (
	"fmt"
	"io"
)

// T :形参的类型，类型
// Number/any  :形参的约束
// func Sum[T any](vals []T) T {
func Sum[T Number](vals []T) T {
	var res T
	println("不能使用generics，这个vals可能传进来的是struct无法计算")
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

// 过滤查找
func Find[T any](vals []T, filter func(t T) bool) []T {
	var result []T
	for _, v := range vals {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
func TryFindAll() {
	// 查找年龄大于等于 30 的人
	filter := func(v int) bool {
		return v >= 30
	}

	result := Find([]int{1, 12, 32, 30, 112}, filter)
	fmt.Println(result) // 输出: {32, 30, 112}
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

// 约束类型
type Number interface {
	//int | uint
	int64 | ~int //int的衍生类型，即Integer等
}

type Integer int

func UseSum() {
	res := Sum[int]([]int{123, 1, 121})
	println(res)
	resV1 := Sum[Integer]([]Integer{122, 11})
	println(resV1)
	println(Sum([]int64{
		12, 1, 11,
	}))
}

func Closable[T io.Closer]() {
	println("将任何实现了 io.Closer 接口的类型作为参数传递，并调用它的 Close 方法。")
	var t T
	t.Close()
}
