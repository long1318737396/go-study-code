package main
import "fmt"
//变量定义
var expA = 1 //全局变量
func main() { //主函数
	var expA int = 2 //局部变量
	fmt.Println(expA, expB)
}

var expB = 5 //全局变量
