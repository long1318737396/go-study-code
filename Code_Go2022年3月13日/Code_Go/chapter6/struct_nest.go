package main
import (
	"fmt"
)

//猫结构体
type Cat struct {
	//名字
	name  string
	//身体数据
	bodyInfo BodyInfo
}
//身体数据结构体
type BodyInfo struct{
	//体重
	weight float64
	//颜色
	color string
}
func main(){
	catOne:=Cat{
		name: "三酷猫",
		bodyInfo: BodyInfo{
			weight: 10.5,
			color: "白色"}}
	fmt.Println("我的名字是：",catOne.name,"，体重：",catOne.bodyInfo.weight,"，毛色：",catOne.bodyInfo.color)
}