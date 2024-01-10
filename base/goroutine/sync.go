package main

import (
	"sync"
	"time"
)

var (
	testMap = make(map[int]int)
	lock    sync.Mutex
)

func testNum(num int) {
	lock.Lock()
	res := 1
	for i := 0; i <= num; i++ {
		res *= 1
	}
	time.Sleep(time.Second)
	testMap[num] = res
	lock.Unlock()
}

func main() {
	//start := time.Now()
	for i := 0; i < 20; i++ {
		go testNum(i)
	}
	time.Sleep(time.Second * 5)

}
