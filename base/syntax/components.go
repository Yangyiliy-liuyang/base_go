package syntax

import "fmt"

type Inner struct {
}

type Outer struct {
	Inner Inner
}

type OuterPtr struct {
	*Inner
}

func (i Inner) Run() {
	fmt.Println("Run By Inner...")
}

func Components() {
	var outer Outer
	outer.Inner.Run()
	//匿名字段可以

	var ErrouterPtr OuterPtr
	ErrouterPtr.Run() //panic: runtime error: invalid memory address or nil pointer dereference
	// 正确写法
	outerPtr := OuterPtr{
		Inner: &Inner{},
	}
	outerPtr.Run()
}
