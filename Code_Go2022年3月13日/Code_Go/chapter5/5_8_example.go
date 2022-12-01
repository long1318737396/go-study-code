package main
import (
	"fmt"
)
func main(){
	var width float64=12
	var height float64=14
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println(err)
		}
	}()
	if width<0||height<0 {
		panic("数值应为正数。")
	}
	fmt.Println("面积为：",width*height)
}
