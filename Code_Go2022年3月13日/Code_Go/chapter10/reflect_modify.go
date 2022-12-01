package main

import (
	"fmt"
	"reflect"
)
func main() {
	numA := 100
	//取numA的地址
	addrValueOfNumA := reflect.ValueOf(&numA)
	fmt.Println(&numA)
	//获取地址所表示的值，即numA的值
	valueOfNumA := addrValueOfNumA.Elem()
	//获取valueOfNumA所在地址
	fmt.Println(valueOfNumA.Addr())
	//判断该值是否可被寻址
	fmt.Println(valueOfNumA.CanAddr())
	//修改numA的值为200
	valueOfNumA.SetInt(200)
	//输出numA变量的值
	fmt.Println(numA)

	//定义结构体
	type cat struct {
		Name string
		age  int
	}
	//声明变量catA，类型为cat，并赋初始值
	catA := &cat{Name: "姓名", age: 18}
	//输出catA的值
	fmt.Println(catA)
	//取catA的地址
	addrValueOfCat := reflect.ValueOf(catA)
	//通过addrValueOfCat变量获取catA的值
	valueOfCat := addrValueOfCat.Elem()
	//找到名为Name的成员变量
	catName := valueOfCat.FieldByName("Name")
	//修改名为Name的成员变量值
	catName.SetString("三酷猫")
	//输出catA的值
	fmt.Println(catA)
}
