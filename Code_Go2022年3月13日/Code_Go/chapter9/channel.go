package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intChan chan int
	var stringChan chan string
	var arrayChan chan []bool
	fmt.Println(intChan,reflect.TypeOf(intChan))
	fmt.Println(stringChan,reflect.TypeOf(stringChan))
	fmt.Println(arrayChan,reflect.TypeOf(arrayChan))	
}