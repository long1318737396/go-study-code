package main
import (
	"fmt"
)
/*位运算符的使用
*/
func main() {
	var numOne int=0                //声明numOne变量
	var numTwo int=1                //声明numTwo变量
	fmt.Println(numOne&numTwo)     //输出numOne和numTwo的按位与结果
	fmt.Println(numOne|numTwo)      //输出numOne和numTwo的按位或结果
	fmt.Println(numOne^numTwo)     //输出numOne和numTwo的按位异或结果
fmt.Println(numOne&^numTwo)    //输出numOne和numTwo的按位清空结果
	var numThree int=20             //声明numThree变量
	fmt.Println(numThree<<2)         //输出numThree左移2位后的结果
	fmt.Println(numThree>>2)         //输出numThree右移2位后的结果
}
