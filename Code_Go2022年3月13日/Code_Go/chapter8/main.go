package main

import (
	"chapter8/model"
	"fmt"
)

func main() {
	personTest := model.NewPerson("三酷猫", 18)
	fmt.Println(personTest.Name)
	personTest.GrowUp()
}
