package main
import (
	"fmt"
)
/*结构体类型
*/

func main() {
	//结构体Person
	type Person struct {
		//string型字段name
		name string
		//int型字段age
		age int
		//int型字段gender，0表示男，1表示女
		gender int
	}
	//实例化Person类型变量alice
	var alice Person
	alice.name="alice"
	alice.gender=1
	alice.age=25
	//输出alice各字段的值
	fmt.Println("姓名：",alice.name)
	fmt.Println("性别：",alice.gender)
	fmt.Println("年龄：",alice.age)
	fmt.Println("alice的值为：",alice)
	//实例化Person类型变量bob
	var bob Person
	bob.name="bob"
	bob.gender=0
	bob.age=28
	//输出bob各字段的值
	fmt.Println("bob的值为：",bob)
	//实例化Person类型变量cindy
	var cindy Person
	cindy.name="alice"
	cindy.gender=1
	cindy.age=18
	//输出cindy各字段的值
	fmt.Println("cindy的值为：",cindy)
}
