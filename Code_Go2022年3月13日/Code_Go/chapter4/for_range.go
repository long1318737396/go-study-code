package main
import (
	"fmt"
)

func main() {
	for key, value := range []string{"你好", "三酷猫"} {
		fmt.Println("key=", key, "，value=", value)
	}
	
	fruit := map[string]float64{
		"apple": 5.4,
		"banana": 6.8,
		"cherry":10.5,
	}
	for key, value := range fruit {
		fmt.Println(key, value)
	}
	
	var str = "Hello，三酷猫"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}
}