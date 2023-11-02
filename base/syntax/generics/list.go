package main

type List[T any] interface {
	ADD(index int, val T)
	Append(val T) error
	Delete(index int) error
}
