package main
import (
	"fmt"
)

func main() {
	var str string="hello，三酷猫"
	switch {
	case str=="hello，三酷猫":
		fmt.Println("英语打招呼",str)
		fallthrough
	case str!="你好，三酷猫":
		fmt.Println("汉语打招呼 你好，三酷猫")
	default:
		fmt.Println("没有看到三酷猫")
	}	
}