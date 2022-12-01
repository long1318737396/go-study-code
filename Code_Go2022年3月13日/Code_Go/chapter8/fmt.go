package main

import (
	"fmt"
	"os"
)

func main() {
	catModel := struct{
		catName string
		age int}{"三酷猫",18}
	fmt.Printf("%v\n", catModel)
	fmt.Printf("%+v\n", catModel)
	fmt.Printf("%#v\n", catModel)
	fmt.Printf("%T\n", catModel)	

	intNum := 26
	fmt.Printf("%b\n", intNum)
	fmt.Printf("%d\n", intNum)
	fmt.Printf("%o\n", intNum)
	fmt.Printf("%x\n", intNum)
	fmt.Printf("%X\n", intNum)
	fmt.Printf("%U\n", intNum)
	fmt.Printf("%q\n", intNum)

	floatNum:=123.456
	fmt.Printf("%b\n", floatNum)
	fmt.Printf("%e\n", floatNum)
	fmt.Printf("%E\n", floatNum)
	fmt.Printf("%f\n", floatNum)
	fmt.Printf("%g\n", floatNum)
	fmt.Printf("%G\n", floatNum)

	strVal :="你好，三酷猫"
	fmt.Printf("%s\n", strVal)
	fmt.Printf("%q\n", strVal)
	fmt.Printf("%x\n", strVal)
	fmt.Printf("%X\n", strVal)

	boolVal :=true
	fmt.Printf("%t\n", boolVal)
	fmt.Println(boolVal)

	strVal ="你好，三酷猫"
	fmt.Printf("%p\n", &strVal)

	numVal :=1234.5678
	fmt.Printf("%f\n", numVal)
	fmt.Printf("%4.2f\n", numVal)
	fmt.Printf("%2.2f\n", numVal)
	fmt.Printf("%.2f\n", numVal)
	fmt.Printf("%7.f\n", numVal)
	fmt.Printf("%7.3f\n", numVal)
	fmt.Printf("%8.4f\n", numVal)

	fmt.Printf("%8.f\n", numVal)
	fmt.Printf("%08.f\n", numVal)
	fmt.Printf("%-8.f\n", numVal)
	fmt.Printf("%#p\n", &numVal)


	fmt.Print("你好，")
	catName := "三酷猫"
	fmt.Printf("%s\n", catName)
	fmt.Println("你好，三酷猫")

	fileObj, err := os.OpenFile("./text.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错：", err)
		return
	}
	textContent := "你好，三酷猫"
	fmt.Fprint(fileObj, textContent)


	strOne := fmt.Sprint("你好")
	strTwo := fmt.Sprintf("，%s","三酷猫！")
	fmt.Println(strOne, strTwo)

	err_ := fmt.Errorf("这个操作包含错误")
	fmt.Println(err_)
	//输出err变量的类型
	fmt.Printf("%T",err_)

	var strOne_ string
	var strTwo_ string
	num,err:=fmt.Scan(&strOne_,&strTwo_)
	fmt.Println(num,err)
	fmt.Printf("输入的字符串：%s %s", strOne_, strTwo_)

}
