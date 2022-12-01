package main
import (
	"fmt"
)
/*转义字符
*/
func main() {
	var text string                  //标准格式声明变量text
	text="唐\t宋\t元\n明\t清\t"     //给变量赋值
	fmt.Println(text)                //输出变量的值
}
