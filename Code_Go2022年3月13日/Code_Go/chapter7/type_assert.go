package main
import (
	"fmt"
)
type EmptyInterface interface {
}
func main() {
	var emptyInterface EmptyInterface
	str:="我是三酷猫"
	emptyInterface=str
	val, boolVal :=emptyInterface.(string)
	if boolVal {
		fmt.Println("emptyInterface保存了字符串类型数据：",val)
	}else{
		fmt.Println("emptyInterface保存的不是字符串类型数据。")
	}
}
