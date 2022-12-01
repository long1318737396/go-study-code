package main
import "fmt"

//叫动作接口，所有动物都应实现该接口
type Sayer interface {
	//叫动作
	say()
}
//跑动作接口，所有动物都应实现该接口
type Runner interface {
	//跑动作
	run()
}
//睡觉动作接口，所有动物都应实现该接口
type Sleeper interface {
	//睡觉动作
	sleep()
}
//定义猫类型
type Cat struct{}
//猫的动作：叫
func (catInstance Cat) say() {
	fmt.Println("喵喵喵")
}
//猫的动作：跑
func (catInstance Cat) run() {
	fmt.Println("奔跑中")
}
//猫的动作：睡觉
func (catInstance Cat) sleep() {
	fmt.Println("睡眠中")
}
func main() {
	var anyAnimalSayer Sayer
	var anyAnimalRunner Runner
	var anyAnimalSleeper Sleeper
	catOne := Cat{}
	anyAnimalSayer=&catOne
	anyAnimalRunner=&catOne
	anyAnimalSleeper=&catOne
	anyAnimalSayer.say()
	anyAnimalRunner.run()
	anyAnimalSleeper.sleep()
}
