package main
import (
	"fmt"
)
/*关系运算符的使用
*/
func main() {
	fmt.Println(100==(50+50))      //输出100与50+50的值是否相等
	fmt.Println((51+49)!=(50*2))    //输出51+49与50*2的值是否不相等
	var text string = "abcde"        //声明字符串类型变量text，值为"abcde"
	fmt.Println(text[0]==97)        //输出text字符串中首个字符的ASCII码是否与97相等
}
