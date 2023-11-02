package main

import "fmt"

func main() {
	f := func1 //方法赋值给变量
	f(1, 2)
	_, _ = func1(1, 2) //方法调用返回赋值
	f2 := func2()      //返回值是方法
	s := f2("张三")
	println("---", s)
	println(f2("lisi"))
	func3()
	func4()
	Func5Invote()
	println("defer和闭包")
	DeferClosure()
	DeferClosureV1()
	println("DeferRutern:", DeferReturn())
	println("DeferRuternV1:", DeferReturnV1())
	println("DeferRuternV2:", DeferReturnV2().name)
	MapDelete()
	println("-----------DeferClosureLoop:--------")
	DeferClosureLoopV1()
	println()
	DeferClosureLoopV2()
	println()
	DeferClosureLoopV3()
}
func func1(a, b int) (int, int) {
	println("func1")
	return a, b
}
func func3() {
	f3 := func(name string) string {
		return name + ",你好"
	}
	s := f3("李白")
	println(s)
}
func func2() func(name string) string {
	return func(name string) string {
		return name + ",你好"
	}
}
func func4() {
	f5 := func(name string) string {
		return name + ",你好"
	}("大名") //立即发起调用
	println("----------")
	println(f5)
	println(func(name string) string {
		return name + ",你好"
	}("大名"))
}

func func5(name string) func() string {
	//闭包：下面的方法和它的上下文 name
	return func() string {
		return name + ",你好"
	}
}
func Func5Invote() {
	println(func5("周达")())
	f := func5("周达") //方法
	println("---", f)
	println(f())
}
func DeferClosure() {
	i := 1
	println("作为闭包引入的，执行defer对应的方法的时候才确定")
	defer func() {
		println(i)
	}()
	i = 1
}
func DeferClosureV1() {
	i := 1
	println("作为参数传入的，定义defer时已经确定")
	defer func(i int) { //参数传递
		println(i)
	}(i)
	i = 1
}
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}
func DeferReturnV1() (a int) {
	println("是带名字的返回值，可以修改这个返回值，否则不能修改。")
	a = 0
	defer func() {
		a = 1
	}()
	return a
}
func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		name: "Tom",
	}
	defer func() {
		res.name = "jack"
	}()
	return res
}

type MyStruct struct {
	name string
}

func MapDelete() {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Printf("%v\n", m)
	for key, value := range m {
		println(key, value)
	}
	delete(m, "key1")
	fmt.Printf("%v\n", m)
	println("没有是否删除成功的标记")
	delete(m, "key3")
	fmt.Printf("%v\n", m)

}
func DeferClosureLoopV1() {
	println("在内存中，迭代参数的位置是一开始就确定了，这个位置的i最后是10")
	for i := 0; i < 10; i++ {
		defer func() {
			print(i, " ")
		}()
	}
}
func DeferClosureLoopV2() {
	println("defer内部是栈")
	for i := 0; i < 10; i++ {
		defer func(val int) {
			print(val, " ")
		}(i)
	}
}
func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			print(j, " ")
		}()
	}
}
