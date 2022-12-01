package main

import (
	"fmt"
	"runtime"
)

func main() {
	go strOutput("三酷猫")
	runtime.Gosched()
	fmt.Println("你好")
}
func strOutput(str string) {
	fmt.Println(str)
}