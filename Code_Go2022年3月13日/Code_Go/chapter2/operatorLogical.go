package main
import (
	"fmt"
)
/*逻辑运算符的使用
*/
func main() {
	fmt.Println(true&&false)      //输出true和false的逻辑与结果
	fmt.Println(true||false)        //输出true和false的逻辑或结果
	fmt.Println(!(true&&false))    //输出true和false的逻辑与后的结果的逻辑非的结果
}
