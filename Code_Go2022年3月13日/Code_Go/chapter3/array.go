package main
import (
	"fmt"
	"reflect"
)
/*数组类型
*/
func main() {
	//通过var关键字声明数组arrZeroInt
	//仅声明数组，编译器会自动初始化数组的元素为数据项类型的零值
	var arrZeroInt = [4]int32{}
	fmt.Println("arrZeroInt=", arrZeroInt)
	//对数组变量使用Go语言内置的len()函数来获取数组长度
	fmt.Println("arrZeroInt的长度为", len(arrZeroInt))

	//通过指定“元素索引值:元素值”
	arrString := [6]string{0: "张三", 3: "李四"}
	fmt.Println("arrString=", arrString)
	fmt.Println("arrString的长度为", len(arrString))

	//指定数组第一个元素为99,第6个元素即索引值为5的元素值为128。
	arrInt64 := [7]int64{99, 5: 128}
	fmt.Println("arrInt64=", arrInt64)
	fmt.Println("arrInt64的长度为", len(arrInt64))

	//与前例类似，但是数组长度由编译器自动推导，数组长度就是给出的最大索引值+1
	arrAutoInit := [...]int{10, 5: 100} //指定索引位置初始化，数组长度与此有关 {10,0,0,0,0,100}
	fmt.Println("arrAutoInit=", arrAutoInit)
	fmt.Println("arrAutoInit的长度为", len(arrAutoInit))

	//通过new关键字声明一个整型数组，此时的arrPointer的数据类型为数组指针
	arrPointer := new([20]int)
	fmt.Println("arrPointer=", arrPointer)
	fmt.Println("arrPointer的第一个元素值为", arrPointer[0])
	fmt.Println("arrPointer的长度为", len(arrPointer))
	//通过标准库reflect包中的reflect.TypeOf函数获取arrPointer的数据类型
	fmt.Println("arrPointer的数据类型为", reflect.TypeOf(arrPointer))

	//函数体内使用:=赋值运算符声明int型数组arrInt
	//声明长度为6，数据项类型为int的数组arrInt并初始化
	arrInt := [6]int{11, 15, 25, 23, 19, 78}
	fmt.Println("arrInt=", arrInt)
	fmt.Println("arrInt的长度为", len(arrInt))
	//获取arrInt的第一个元素
	fmt.Println("arrInt的第一个元素为", arrInt[0])
	//可以通过Go语言标准库的reflect包中的TypeOf函数查看数组arrInt的数据类型
	fmt.Println("arrInt的数据类型为", reflect.TypeOf(arrInt))

	//定义长度由编译器自动推导，数据项类型为float64的arrFloat64
	arrAutoFloat64 := [...]float64{1.31, 3.14, 5.28, 6.78}
	fmt.Println("arrAutoFloat64=", arrAutoFloat64)
	fmt.Println("arrAutoFloat64的长度为", len(arrAutoFloat64))

	//多维数组
	arrMulti := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arrMulti=", arrMulti)
	fmt.Println("arrMulti的长度为", len(arrMulti))

	fmt.Println("arrMulti[0][1]=", arrMulti[0][1])

	arrMulti[0][1]=100
	fmt.Println("arrMulti=", arrMulti)

	fmt.Println("arrMulti的数据类型为", reflect.TypeOf(arrMulti))
}
