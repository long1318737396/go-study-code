package main

import (
	"fmt"
	"sync"
)

var syncWait sync.WaitGroup
func main() {
	syncWait.Add(1)
	go func() {
		defer syncWait.Done()
		fmt.Println("你好，三酷猫")
	}()
	syncWait.Wait()
	fmt.Println("程序运行结束")
}