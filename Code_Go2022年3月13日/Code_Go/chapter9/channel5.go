package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)
	go recvFunc(chan1,chan2)
	go sendFunc1(chan1)
	go sendFunc2(chan2)
	time.Sleep(5 * time.Second)
	fmt.Println("main()函数结束")
}
func sendFunc1(chan1 chan int){
	for i := 0; i < 5; i++ {
		chan1 <- i
		time.Sleep(1 * time.Second)
	}
}
func sendFunc2(chan2 chan int){
	for i := 10; i >= 5; i-- {
		chan2 <- i
		time.Sleep(1 * time.Second)
	}
}
func recvFunc(chan1 chan int,chan2 chan int){
	for {
		select {
		case intValue1 := <-chan1:
			fmt.Println("接收到chan1通道的值：", intValue1)
		case intValue2 := <-chan2:
			fmt.Println("接收到chan2通道的值：", intValue2)
		}
	}
}