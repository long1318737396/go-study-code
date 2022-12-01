package main
import "fmt"
func main() {
	//单一常量定义使用
	const age = 1                           //定义值为1的整型常量age
	const num int = 10                      //定义值为10的整型常量num，int指type
	fmt.Println("三酷猫!", age, "岁", num, "只") //打印输出常量，通过逗号分割可以输出多个常量

	//批量常量声明及使用
	const (
		e  = 2.7182818 //数学中的自然常数e
		pi = 3.1415926 //数学中的七位小数圆周率π
	)
	fmt.Println("数学里的常量：", e, pi) //打印输出常量

	//常量生成器iota
	const (
		spring = iota //初始值为0
		summer
		autumn
		winter
	)
	fmt.Println("一年四季：", spring, summer, autumn, winter)

}
