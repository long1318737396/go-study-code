package main

import (
	"strconv"
	"fmt"
)

func main() {
	strVal:="200"
	//将字符串型数据转为整数型
	intVal,err:=strconv.Atoi(strVal)
	if err!=nil{
		fmt.Println("转换整型出错，错误信息：",err)
	}
	fmt.Println(intVal)
	fmt.Printf("%T\n",intVal)
	//将整数型数据转为字符串型
	strVal2:=strconv.Itoa(intVal)
	fmt.Println(strVal2)
	fmt.Printf("%T\n",strVal2)	

	boolVal, err := strconv.ParseBool("true")
	if err!=nil{
		fmt.Println("转换出错，错误信息：",err)
		return
	}
	fmt.Println(boolVal)
	fmt.Printf("%T\n",boolVal)
	floatVal, err := strconv.ParseFloat("123.987", 64)
	if err!=nil{
		fmt.Println("转换出错，错误信息：",err)
		return
	}
	fmt.Println(floatVal)
	fmt.Printf("%T\n",floatVal)
	intVal_, err := strconv.ParseInt("-500", 10, 64)
	if err!=nil{
		fmt.Println("转换出错，错误信息：",err)
		return
	}
	fmt.Println(intVal_)
	fmt.Printf("%T\n",intVal_)
	uintVal, err := strconv.ParseUint("500", 10, 64)
	if err!=nil{
		fmt.Println("转换出错，错误信息：",err)
		return
	}
	fmt.Println(uintVal)
	fmt.Printf("%T\n",uintVal)

	strVal1:= strconv.FormatBool(true)
	fmt.Println(strVal1)
	fmt.Printf("%T\n",strVal1)
	strVal2= strconv.FormatFloat(123.9876,'E',-1,64)
	fmt.Println(strVal2)
	fmt.Printf("%T\n",strVal2)
	strVal3 := strconv.FormatInt(-500,10)
	fmt.Println(strVal3)
	fmt.Printf("%T\n",strVal3)
	strVal4:= strconv.FormatUint(500,10)
	fmt.Println(strVal4)
	fmt.Printf("%T\n",strVal4)
}
