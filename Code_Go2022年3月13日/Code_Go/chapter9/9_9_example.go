package main

import (
	"fmt"
)

func main() {
	fmt.Println(SievePrime(1000))
}
func SievePrime(endNum int) []int {
	//声明变量numArray，下标代表数字，值表示是否为素数，是素数为0，反之为1
	numArray := make([]int, endNum+1)
	//0和1均不是素数，标记为1
	numArray[0], numArray[1] = 1, 1
	for i := 2; i <= endNum; i++ {
		if numArray[i] == 0 {
			for j := 2 * i; j <= endNum; j += i {
				numArray[j] = 1
			}
		}
	}
	result := make([]int, 0)
	for i := 2; i <= endNum; i++ {
		if numArray[i] == 0 {
			//将标记为0（是素数）的下标值添加到result中
			result = append(result, i)
		}
	}
	return result
}