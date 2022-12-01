package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	close(intChan) 
	for {
		intValue, isOpen := <-intChan
		if !isOpen { 
			fmt.Println("通道已关闭")
			break
		}
		fmt.Println(intValue)
	}
}
