package main
import (
	"fmt"
)
/*解答三酷猫关于指针的困惑
*/
func main() {
	var c int=100
	var b *int=&c
	var a **int=&b
	var d ***int=&a
	fmt.Println(b==*a)
	fmt.Println(**d==b)
	fmt.Println(*d==a)
}
