package main
import "fmt"
/*指针运算符的使用
*/
func main() {
	var numOne int=5                //声明numOne变量，类型为int，值为5
	var pointer *int=&numOne         //声明pointer变量，类型为指针变量，值为numOne变量的内存地址
	fmt.Println(&numOne)            //输出numOne变量的实际内存地址
	fmt.Println(*pointer)              //输出pointer变量表示的内存地址的变量的值
}
