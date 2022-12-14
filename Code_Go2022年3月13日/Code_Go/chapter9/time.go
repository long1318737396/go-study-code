package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	fmt.Println(time.Now())
	<-time.After(2*time.Second)
	fmt.Println(time.Now())

	fmt.Println(time.Now())
	time.AfterFunc(2*time.Second,sayHello)
	time.Sleep(3*time.Second)
}	

func sayHello(){
	fmt.Println("三酷猫打招呼",time.Now())
}