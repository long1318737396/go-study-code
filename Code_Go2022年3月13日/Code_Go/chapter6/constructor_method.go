package main
import (
	"fmt"
	"reflect"
)
type Cat struct {
	//性别，0为男，1为女
	gender int
	//颜色
	color string
	//名字
	name  string
}
//构造函数
func newCat(gender int,color string,name string) *Cat {
	return &Cat{
		gender: gender,
		color: color,
		name: name}
}
//吃饭
func (catInstance Cat) eat(food string){
	fmt.Println(catInstance.name,"正在吃：",food)
}
//睡觉
func (catInstance Cat) dream(){
	fmt.Println(catInstance.name,"睡得正香")
}
//喵喵叫
func (catInstance Cat) mewing(){
	fmt.Println(catInstance.name,"喵喵喵")
}
//改名
func (catInstance *Cat) rename(newName string){
	catInstance.name=newName
	fmt.Println("方法中catInstance的name属性值为",catInstance.name)
}

type MyString string
func(name MyString) whoAmI(){
	fmt.Println("我的名字是：", name)
}


func main(){
	//通过newCat()函数生成Cat结构
	catOne:=newCat(0,"white","三酷猫")
	fmt.Println(catOne)
	fmt.Println(reflect.TypeOf(catOne))
	//catOne执行吃饭动作
	catOne.eat("鱼")
	//catOne执行睡觉动作
	catOne.dream()
	//catOne执行喵喵叫动作
	catOne.mewing()
	//修改catOne的名字
	catOne.rename("另一只三酷猫")
	fmt.Println(catOne)

	var me MyString="三酷猫"
	me.whoAmI()
}
