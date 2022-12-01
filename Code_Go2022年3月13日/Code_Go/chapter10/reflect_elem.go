package main

import (
	"fmt"
	"reflect"
)
func main() {
    intNum := 200
    ins := &intNum
    //获取指针变量的类型名和种类
    typeOfIntNumPtr := reflect.TypeOf(ins)
    fmt.Println(typeOfIntNumPtr.Name(), typeOfIntNumPtr.Kind())
    //获取指针变量所表示的变量的类型名和种类
    typeOfIntNumPtr = typeOfIntNumPtr.Elem()
    fmt.Println(typeOfIntNumPtr.Name(), typeOfIntNumPtr.Kind())
}
