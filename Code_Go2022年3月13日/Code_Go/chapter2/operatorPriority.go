package main
import (
	"fmt"
)
/*运算符的优先级
*/
func main() {
	var numA int=10                //声明numA变量
	var numB int=20                //声明numB变量
	var numC int=30                //声明numC变量
	var numD int=40                //声明numD变量
	fmt.Println((numA+numB)*numC+numD)
	fmt.Println(numA+numB*(numC+numD))
	fmt.Println(numA+numD==numB+numC)
	fmt.Println(numD-numC==numA+numB)
	fmt.Println(numD-numC==(numA+numB))
	fmt.Println(numC==numA+numB)
}
