package main
import (
	"fmt"
)
/*复数型变量
*/
func main() {
	var complexOne complex128                  //标准格式声明变量complexOne
	complexOne=complex(5,10)                  //给变量赋值
	fmt.Println(complexOne)                     //打印输出复数的值
	fmt.Println(real(complexOne))                 //打印输出实部
	fmt.Println(imag(complexOne))                //打印输出虚部
}
