package main

import (
	"fmt"
	"time"
)

var count int
func main() {
	go countPlus(10000)
	go countPlus(10000)
	time.Sleep(2*time.Second)
	fmt.Println(count)
}
func countPlus(times int){
	for i:=0;i<times;i++{
		count++
	}
}