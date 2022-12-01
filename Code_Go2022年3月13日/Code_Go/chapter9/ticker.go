package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		sayHello()
	}
}
func sayHello(){
	fmt.Println("三酷猫打招呼",time.Now())
}