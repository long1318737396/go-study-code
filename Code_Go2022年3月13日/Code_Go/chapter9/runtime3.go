package main

import (
	"fmt"
	"runtime"
)

func main() {
	go loop(100)
	runtime.Gosched()
	fmt.Println("main()函数执行结束")
}
func loop(maxTime int) {
	for i:=0;i<maxTime;i++{
		if i>=3{
			runtime.Goexit()
		}
		fmt.Println("当前循环次数：",i)
	}
}
