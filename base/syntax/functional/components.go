package main

func main() {
	components()
	components1()
}

type Inter struct {
}
type Outer struct {
	Inter
}
type Outer1 struct {
	//Inter
	*Inter
}

func (i Inter) SayHi() {
	println("hi")
}

func components() {
	var o Outer
	o.SayHi()
}
func components1() {
	println("内部的指针需要初始化：")
	o1 := &Outer1{
		Inter: &Inter{},
	}
	o1.SayHi()
}
