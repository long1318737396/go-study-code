package main
import (
	"fmt"
)
/*算术运算符的使用
*/
func main() {
	var expNumOne=5                      //声明expNumOne变量
	var expNumTwo=6                      //声明expNumTwo变量
	fmt.Println(expNumOne+expNumTwo)      //输出两个数相加的结果
	fmt.Println(expNumOne-expNumTwo)      //输出两个数相减的结果
	fmt.Println(expNumOne*expNumTwo)      //输出两个数相乘的结果
	fmt.Println(expNumTwo/expNumOne)      //输出两个数相除的结果
	fmt.Println(expNumTwo%expNumOne)     //输出两个数相除取余数的结果
	expNumOne++                         //expNumOne变量自加1
	fmt.Println(expNumOne)                 //输出运算结果
	expNumTwo--                         //expNumTwo变量自减1
	fmt.Println(expNumTwo)                //输出运算结果
	var uInt8Max uint8=255                 //声明uInt8Max变量，类型为uint8，值为该类型最大值
	fmt.Println(uInt8Max+1)                //输出运算结果
	var int8Max int8=127                 //声明int8Max变量，类型为uint8，值为该类型最大值
	fmt.Println(int8Max+1)                //输出运算结果
}
