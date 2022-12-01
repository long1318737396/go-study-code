package main

import (
	"fmt"
	"time"
	"sync"
)

var count int
var locker sync.RWMutex
func main() {
	for i := 1; i <= 3; i++ {
		go write(i)
	}
	for i := 1; i <= 3; i++ {
		go read(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("count值为：", count)
}
func read(i int) {
	fmt.Println("读操作",i)
	locker.RLock()
	fmt.Println(i,"读count的值为",count)
	time.Sleep(1 * time.Second)
	locker.RUnlock()
}
func write(i int) {
	fmt.Println("写操作",i)
	locker.Lock()
	count++
	fmt.Println(i,"写count的值为",count)
	time.Sleep(1 * time.Second)
	locker.Unlock()
}