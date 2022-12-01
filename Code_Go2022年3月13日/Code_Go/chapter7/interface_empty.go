package main
import (
	"fmt"
	"reflect"
)

type EmptyInterface interface {
}

func main() {
	var emptyInterface EmptyInterface
	//使用空接口保存字符型变量
	strVar:="我是三酷猫"
	emptyInterface= strVar
	fmt.Println(emptyInterface)
	fmt.Println(reflect.TypeOf(emptyInterface))
	//使用空接口保存数字型变量
	numVar:=18
	emptyInterface= numVar
	fmt.Println(emptyInterface)
	fmt.Println(reflect.TypeOf(emptyInterface))
	//使用空接口保存布尔型变量
	boolVar:=true
	emptyInterface=boolVar
	fmt.Println(emptyInterface)
	fmt.Println(reflect.TypeOf(emptyInterface))

	var animalInfo=make(map[string]EmptyInterface)
	animalInfo["name"]="三酷猫"
	animalInfo["age"]=3
	animalInfo["married"]=false
	fmt.Println(animalInfo)
}

