package main

import (
	"fmt"
	"sync"
)

var syncWait sync.WaitGroup
func main() {
	syncWait.Add(1)
	go hello()
	fmt.Println("程序运行结束")
	syncWait.Wait()
}
func hello(){
	defer syncWait.Done()
	fmt.Println("你好，三酷猫")
}