package main

import (
	"fmt"
	"sync"
)

var syncWait sync.WaitGroup
func main() {
	for i := 0; i < 10; i++ {
		syncWait.Add(1)
		go hello(i)
	}
	syncWait.Wait()
	fmt.Println("程序运行结束")
}
func hello(index int)  {
	defer syncWait.Done()
	fmt.Println("你好，三酷猫",index)
}