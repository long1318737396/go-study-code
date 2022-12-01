package main
import (
	"fmt"
)
/*截取特定范围的字符串
*/
func main() {
	var text string                    //标准格式声明变量text
	text="Hello，三酷猫"             //给变量赋值
	fmt.Println(text[8:])               //打印输出最后三个中文字符
}
