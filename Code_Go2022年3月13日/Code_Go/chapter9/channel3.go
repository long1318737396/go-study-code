package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//声明int型通道变量intChan，缓冲区大小为2
	var intChan=make(chan int,2)
	//输出intChan变量的值和类型
	fmt.Println(intChan,reflect.TypeOf(intChan),len(intChan),cap(intChan))
	//将整型值100通过intChan通道发送
	intChan<-100
	fmt.Println(len(intChan),cap(intChan))
	intChan<-200
	fmt.Println(len(intChan),cap(intChan))
	go intChanRecvListener(intChan)
	time.Sleep(2*time.Second)
	//关闭intChan通道
	close(intChan)
}
func intChanRecvListener(intChan chan int){
	//读取intChan通道中的值，并将结果赋值给intValue变量
	intValue:=<-intChan
	//输出intValue变量的值和类型
	fmt.Println(intValue,reflect.TypeOf(intValue))
	fmt.Println("成功接收第1个值",len(intChan),cap(intChan))
	intValue=<-intChan
	//输出intValue变量的值和类型
	fmt.Println(intValue,reflect.TypeOf(intValue))
	fmt.Println("成功接收第2个值",len(intChan),cap(intChan))
}