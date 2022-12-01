package main

import (
	"fmt"
	"reflect"
)
func main() {
	//将函数包装为反射值变量
	sayHelloValue := reflect.ValueOf(sayHello)
	//定义函数传入的参数切片
	sayHelloParam := []reflect.Value{reflect.ValueOf("三酷猫"), reflect.ValueOf("你好呀！")}
	//通过反射调用函数
	results := sayHelloValue.Call(sayHelloParam)
	//输出函数执行结果
	fmt.Println(results)
}
func sayHello(name string, content string) string {
	output := name + "说：" + content
	return output
}