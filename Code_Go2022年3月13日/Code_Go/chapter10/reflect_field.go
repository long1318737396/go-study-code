package main

import (
	"fmt"
	"reflect"
)
func main() {
	//声明cat类型结构体
	type cat struct {
		name string `meaning:"全名"`
		age  int   `meaning:"年龄"`
	}
	//声明catOne类型变量并赋初始值
	catOne := cat{name: "三酷猫", age: 18}
	//通过反射获取catOne变量的类型名和种类
	typeOfCat := reflect.TypeOf(catOne)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	//通过反射获取catOne变量的值
	valueOfCat := reflect.ValueOf(catOne)
	//通过反射获取结构体成员的名称和值
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Println(fieldType.Index, fieldType.Name, valueOfCat.Field(i), fieldType.Tag)
	}
	//查找名为age的成员
	catType, ok := typeOfCat.FieldByName("age")
	if ok {
		//输出age成员的部分Tag文本
		fmt.Println(catType.Tag.Get("meaning"))
	}
}