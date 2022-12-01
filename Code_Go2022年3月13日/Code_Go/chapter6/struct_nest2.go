package main
import (
	"fmt"
)
//猫结构体
type Cat struct {
	//眼睛的颜色
	eyeColor string
	//动物结构体
	animal *Animal
}
//猫-喵喵叫
func (catInstance Cat) mewing(){
	fmt.Println(catInstance.animal.name,"喵喵喵")
}
//狗结构体
type Dog struct {
	//身体的颜色
	bodyColor string
	//动物结构体
	animal *Animal
}
//狗-汪汪叫
func (dogInstance Dog) bowwow(){
	fmt.Println(dogInstance.animal.name,"汪汪汪")
}
//动物结构体
type Animal struct{
	//名字
	name  string
}
func main(){
	dogOne:=&Dog{bodyColor: "黑色",animal: &Animal{
		name: "贝贝"}}
	fmt.Println(dogOne.animal.name,"身体的颜色是",dogOne.bodyColor)
	dogOne.bowwow()
	catOne:=&Cat{eyeColor: "蓝色",animal: &Animal{
		name: "三酷猫"}}
	fmt.Println(catOne.animal.name,"眼睛的颜色是",catOne.eyeColor)
	catOne.mewing()
}