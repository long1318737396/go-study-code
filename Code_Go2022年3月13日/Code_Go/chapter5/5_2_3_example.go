package main
import (
	"fmt"
	"reflect"
)
/*三酷猫识别数据类型
 */
func main() {
	//函数测试代码
	getVarType("abcd",0,true,[3]int{0,1,2},[]string{})
}
func getVarType(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(reflect.TypeOf(arg))
	}
}
