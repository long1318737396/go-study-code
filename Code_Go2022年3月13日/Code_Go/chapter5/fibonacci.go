package main
import (
	"fmt"
)
/*输出斐波那契数列
 */
func main() {
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = fibonacci(i)
	}
	fmt.Println(array)
}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = n
	}else{
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
