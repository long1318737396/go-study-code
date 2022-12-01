package main

import (
	"fmt"
	"time"
	"sync"
)

var count int
var locker sync.Mutex
func main() {
	go countPlus(10000)
	go countPlus(10000)
	time.Sleep(2*time.Second)
	fmt.Println(count)
}
func countPlus(times int){
	for i:=0;i<times;i++{
		locker.Lock()
		count++
		locker.Unlock()
	}
}