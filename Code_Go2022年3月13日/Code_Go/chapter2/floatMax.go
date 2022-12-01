package main
import (
	"fmt"
	"math"
)
/*浮点型变量
*/
func main() {
	var float32Number float32                  //标准格式声明变量float32Number
	float32Number=math.MaxFloat32            //给变量赋值
	fmt.Println(float32Number)                 //打印输出
	var float64Number float64                  //标准格式声明变量float64Number
	float64Number=math.MaxFloat64            //给变量赋值
	fmt.Println(float64Number)                 //打印输出
}
