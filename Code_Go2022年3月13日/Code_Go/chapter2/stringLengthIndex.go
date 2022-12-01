package main
import (
	"fmt"
)
/*字符串总长度及特定字符的获取
*/
func main() {
	var text string                  //标准格式声明变量text
	text="Hello，三酷猫"           //给变量赋值
	fmt.Println(len(text))            //打印输出字符串长度
	fmt.Println(text[0])              //打印输出字符串第一个字符
}
