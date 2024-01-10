package main

import "fmt"

var intChan chan int

func main() {
	intChan = make(chan int, 1)
	intChan <- 1
	fmt.Println(<-intChan)
	// 超出范围会导致死锁
	//fmt.Println(<-intChan) //fatal error: all goroutines are asleep - deadlock!

	dogChan1 := make(chan dog, 1)
	dogChan1 <- dog{Name: "小黄"}
	println((<-dogChan1).Name)

	dogChan2 := make(chan interface{}, 10) // 接口
	dogChan2 <- dog{Name: "小黄2"}
	d := (<-dogChan2).(dog) // 使用类型断言
	println(d.Name)
	dogChan2 <- "11111"
	dogChan2 <- 11111
	close(dogChan2) // 关闭管道后，无法写入
	for {
		val, ok := <-dogChan2
		if !ok { // 有数据，返回true
			break
		}
		println(val)
	}
}

type dog struct {
	Name string
}
