package main

type I interface {
	M()
}

type T struct {
}

func (T) M() {
}

func main() {
	var t T
	var i I = t
	switch i.(type) {
	case T:
		println("it is type T")
		//case int:
		//	println("it is type int")
		//case string:
		//	println("it is type string")
	}
}
