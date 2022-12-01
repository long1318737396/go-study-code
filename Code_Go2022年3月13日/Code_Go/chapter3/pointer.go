package main
import (
	"fmt"
)
/*指针类型
*/
func main() {
	//指针的一般使用
	var numA int=10                 //声明numA变量，类型为int，值为10
	var intPointer *int                //声明intPointer变量，类型为指针变量
	intPointer=&numA               //获取numA的内存地址，将该地址赋值给intPointer
	fmt.Println(intPointer)             //输出intPointer的值
	fmt.Println(*intPointer)            //通过intPointer的值，获得numA的值，并输出

	//空指针
	var intPointer2 *int                //声明intPointer变量，类型为指针变量
	fmt.Println(intPointer2)             //输出intPointer的值

	//指向指针的指针变量
	var c int=100           //声明整型变量c，值为100
	var b *int=&c           //声明指针型变量b，用来存放整型变量c的地址
	var a **int=&b          //声明指针型变量a，用来存放指针型变量b的地址
	fmt.Println(**a)         //输出a所表示的变量的值所表示的变量的值（b所表示的变量的值，即整型变量c的值）
}
