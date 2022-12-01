package model

import "fmt"

type person struct {
	//小写开头的变量是私有变量，不可以在其它包访问
	age int
	//大写开头的变量是公开变量，可以在其它包访问
	Name string
}

//构造函数，返回person结构体变量的指针
//大写开头的函数是公开函数，可以在其它包访问
func NewPerson(personName string, personAge int) *person {
	return &person{
		Name: personName,
		age:  personAge,
	}
}

//person年龄增长
func (personInstance *person) GrowUp() {
	personInstance.age++
	fmt.Println("年龄增长至", personInstance.age)
}
