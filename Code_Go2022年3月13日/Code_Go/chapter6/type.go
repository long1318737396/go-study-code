package main
import (
	"fmt"
	"reflect"
)

type NewInt = int

type Book struct {
	title string
	author string
	subject string
 } 

func main() {
	type NewInt int
	var intNum NewInt=10
	fmt.Println("intNum的值为：",intNum,"，类型为：",reflect.TypeOf(intNum))

	var bookOne Book
	bookOne.title="书籍名称"
	bookOne.author="作者名称"
	bookOne.subject="书籍主题"
	fmt.Println(bookOne)
	fmt.Println(reflect.TypeOf(bookOne))
	
	bookOne=Book{
		title:"书籍名称",
		author: "作者名称",
		subject: "书籍主题"}
	fmt.Println(bookOne)
	fmt.Println(reflect.TypeOf(bookOne))
	
	bookOne=Book{
		subject: "书籍主题",
		title:"书籍名称",
		author: "作者名称"}
	fmt.Println(bookOne)
	fmt.Println(reflect.TypeOf(bookOne))

	bookOne=Book{
		title: "书籍名称"}
	fmt.Println(bookOne)
	fmt.Println(reflect.TypeOf(bookOne))
	
	bookOne=Book{
		"书籍名称",
		"作者名称",
		"书籍主题"}
	fmt.Println(bookOne)

	testStruct :=struct{
		intA int8
		intB int8
		intC int8
		intD int8}{
		1,
		2,
		3,
		4}
	fmt.Println(&testStruct.intA)
	fmt.Println(&testStruct.intB)
	fmt.Println(&testStruct.intC)
	fmt.Println(&testStruct.intD)

	var bookTwo=new(Book)
	fmt.Println(&bookTwo)
	fmt.Println(bookTwo)
	fmt.Println(reflect.TypeOf(bookTwo))
	bookTwo.author="作者名称"
	bookTwo.title="书籍名称"
	bookTwo.subject="书籍主题"
	fmt.Println(bookTwo)
	
	bookThree:=&Book{}
	fmt.Println(&bookThree)
	fmt.Println(bookThree)
	fmt.Println(reflect.TypeOf(bookThree))
	bookThree.author="作者名称"
	bookThree.title="书籍名称"
	bookThree.subject="书籍主题"
	fmt.Println(bookThree)

}

