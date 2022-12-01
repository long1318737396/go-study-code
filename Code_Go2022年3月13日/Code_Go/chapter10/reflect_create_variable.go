package main

import (
	"fmt"
	"reflect"
)
func main() {
	//声明int型变量numA
	var numA int
	//获取numA的类型，并存入变量typeOfNumA
	typeOfNumA := reflect.TypeOf(numA)
	//根据类型创建变量numAVal
	numAVal := reflect.New(typeOfNumA)
	//输出numAVal的值、类型名称和种类名称
	fmt.Println(numAVal, numAVal.Type(), numAVal.Kind())
	//修改numAVal表示的值
	numAVal.Elem().SetInt(100)
	//输出numAVal表示的值
	fmt.Println(numAVal.Elem().Int())
}
