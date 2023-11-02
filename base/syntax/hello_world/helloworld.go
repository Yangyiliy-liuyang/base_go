package main

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	print("hello go\n")
	Hello()
	String()
	Byte()
	Bool()
	variables()
}
func String() {
	//I say :"hello im liming"
	println("I say :\"hello im liming\"")
	print(`------------------
			hello
go
`)
	println("hello" + strconv.Itoa(123))
	cutPrefix, _ := strings.CutPrefix("aaabbbb", "a")
	println(cutPrefix)
	println(utf8.RuneCountInString("hello中国"))

}
func Byte() {
	var b byte = 'a'
	print("输出的是ASCII码")
	println(b)
}
func Bool() {
	println(true && false)
	println(true || false)
	println(!(true && false)) //! (a && b) === !a || !b
	println(!(true || false)) //! (a || b) === !a && !b

}

var (
	name = "libai"
	Age  = 13
)

const (
	status0 = iota
	status1 = iota
	status2 = iota

	status3 = iota
)

func variables() {
	var i int = 123
	println(i)

}
