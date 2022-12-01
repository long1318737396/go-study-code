package main
import (
	"fmt"
)
/*浮点型简单运算
*/
func main() {
	floatNumOne := 1.0 / 3.0                              //短格式声明变量floatNumOne
	fmt.Println(floatNumOne + floatNumOne + floatNumOne)   //输出floatNumOne相加3次之和
	floatNumTwo := 0.1                                  //短格式声明变量floatNumTwo
	floatNumTwo += 0.2                                 //floatNumTwo与0.2相加
	fmt.Println(floatNumTwo)                             //输出floatNumTwo的值
}
