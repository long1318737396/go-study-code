package main

import (
	"fmt"
	"reflect"
)

func main() {
	//声明int型通道变量intChan
	var intChan=make(chan int)
	//输出intChan变量的值和类型
	fmt.Println(intChan,reflect.TypeOf(intChan))
	go intChanRecvListener(intChan)
	//将整型值100通过intChan通道发送
	intChan<-100
	//关闭intChan通道
	close(intChan)
}
func intChanRecvListener(intChan chan int){
	//读取intChan通道中的值，并将结果赋值给intValue变量
	intValue:=<-intChan
	//输出intValue变量的值和类型
	fmt.Println(intValue,reflect.TypeOf(intValue))
}
