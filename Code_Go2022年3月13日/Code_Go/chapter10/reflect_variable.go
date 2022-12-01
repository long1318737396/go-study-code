package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intNum int
	typeOfIntNum := reflect.TypeOf(intNum)
	fmt.Println(typeOfIntNum.Name(), typeOfIntNum.Kind())

	type structType struct {}
	typeOfStructType:=reflect.TypeOf(structType{})
	fmt.Println(typeOfStructType.Name(), typeOfStructType.Kind())

	intNum = 100
	fmt.Println(reflect.ValueOf(intNum))
	valueOfIntNum := reflect.ValueOf(intNum)
	//类型强制转换为Int
	var originIntNum int64 = int64(valueOfIntNum.Int())
	fmt.Println(originIntNum)

	var intNums = []int{1,3,5,7,9}
	valueOfIntNum = reflect.ValueOf(intNums)
	fmt.Println(valueOfIntNum.IsNil())
	fmt.Println(valueOfIntNum.IsValid())

}