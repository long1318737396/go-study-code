package main
import (
	"fmt"
)

func main() {
	//声明布尔型变量condition，值为true
	var condition bool=true
	//布尔类型判断：condition是否为true
	if(condition){
		fmt.Println("condition的值为true")
	}
	//声明整型变量numA，值为100
	var numA int=100
	//逻辑判断：numA是否大于0
	if numA>0{
		fmt.Println("numA的值大于0")
	}

	if numB:=(10-9);numB>0{
		fmt.Println("计算结果大于0")
	}


	//声明布尔型变量condition，值为true
	condition=true
	//布尔类型判断：condition是否为true
	if(condition){
		fmt.Println("condition的值为true")
	}else{
		fmt.Println("condition的值为false")
	}
	//声明整型变量numA，值为100
	numA=100
	//逻辑判断：numA是否大于0
	if numA>0{
		fmt.Println("numA的值大于0")
	}else{
		fmt.Println("numA的值不大于0")
	}

	var numC=500
	if numC<0{
		fmt.Println("numC的值为负")
	}else if numC==0{
		fmt.Println("numC的值为0")
	}else if numC<100{
		fmt.Println("numC的值小于100")
	}else if numC<200{
		fmt.Println("numC的值小于200")
	}else if numC<300{
		fmt.Println("numC的值小于300")
	}else if numC<400{
		fmt.Println("numC的值小于400")
	}else if numC<500{
		fmt.Println("numC的值小于500")
	}else if numC<600{
		fmt.Println("numC的值小于600")
	}else{
		fmt.Println("numC的值大于或等于1000")
	}


}