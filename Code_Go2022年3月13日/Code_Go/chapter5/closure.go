package main
import (
	"fmt"
)
func main() {
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	accumulator3 := Accumulate(100)
	fmt.Println(accumulator3())
	fmt.Println(accumulator3())
	fmt.Println(accumulator3())
	//再次通过每个变量调用各自的函数
	fmt.Println(accumulator())
	fmt.Println(accumulator2())
	fmt.Println(accumulator3())
}

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
