package main

import (
	"fmt"
)

//从2开始生成整数
func GenNums() chan int {
	num := make(chan int)
	go func() {
		for i := 2; ; i++ {
			num <- i
		}
	}()
	return num
}
//过滤素数
func SievePrime(numCh <-chan int, prime int) chan int {
	result := make(chan int)
	go func() {
		for {
			if i := <-numCh; i%prime != 0 {
				result <- i
			}
		}
	}()
	return result
}
func main() {
	numCh := GenNums()
	for i := 0; i < 100; i++ {
		result := <-numCh
		fmt.Println(result)
		numCh = SievePrime(numCh, result)
	}
}