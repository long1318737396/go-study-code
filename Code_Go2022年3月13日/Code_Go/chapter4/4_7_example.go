package main
import "fmt"
/*三酷猫学算法之冒泡排序
 */
func main() {
	arr := [...]int{100,205,113,15,6,78,24,-10,45,0}
	var n = len(arr)
	for i := 0; i <= n-1; i++ {
		for j := i; j <= n-1; j++ {
			if arr[i] > arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
		}
	}
	fmt.Println(arr)
}