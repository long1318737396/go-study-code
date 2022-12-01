package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取CPU核心数量
	fmt.Println(runtime.NumCPU())
	//设置可用的CPU核心数量（返回之前设置的核心数）
	cpuNum:=runtime.GOMAXPROCS(4)
	fmt.Println(cpuNum)
	cpuNum=runtime.GOMAXPROCS(4)
	fmt.Println(cpuNum)
}