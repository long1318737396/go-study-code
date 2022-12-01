package main
import (
	"fmt"
	"math"
)
/*使运算结果保持特定的精度
*/
func main() {
	var float32Number float32                  //标准格式声明变量float32Number
	float32Number=math.MaxFloat32            //给变量赋值
	fmt.Printf("%9.2e",float32Number)            //打印输出
	fmt.Println()                             //打印空行
	var float64Number float64                  //标准格式声明变量float64Number
	float64Number=math.MaxFloat64            //给变量赋值
	fmt.Printf("%.2e",float64Number)            //打印输出
	fmt.Println()                             //打印空行
	fmt.Printf("%9.2f",123.4567)                 //打印输出

}
