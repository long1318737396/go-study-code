package main
import "fmt"
func main() {
	fmt.Println("defer行为开始")
	defer fmt.Println("操作1")
	defer fmt.Println("操作2")
	defer fmt.Println("操作3")
	fmt.Println("defer行为结束")

	i := 0
	defer fmt.Println(i)
	i++
    fmt.Println(i)
	return

	fmt.Println(funcA())

	numArr := [5]int{1, 2, 3, 4, 5}
	for _, value := range numArr {
		defer fmt.Println(value)
	}
	fmt.Println("函数开始执行")

	numArr2 := [5]int{1, 2, 3, 4, 5}
	defer fmt.Println(numArr2[0])
	defer fmt.Println(numArr2[1])
	defer fmt.Println(numArr2[2])
	defer fmt.Println(numArr2[3])
	defer fmt.Println(numArr2[4])
	fmt.Println("函数开始执行")

	numArr3 := [5]int{1, 2, 3, 4, 5}
	for _, value := range numArr3 {
		defer func() {
			fmt.Println(value)
		}()
	}
	fmt.Println("函数开始执行")

}

func funcA() (numA int) {
	defer func() {
		numA++
	}()
	return 0
}

