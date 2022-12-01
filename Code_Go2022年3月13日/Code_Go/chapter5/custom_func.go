package main
import (
	"fmt"
	"time"
)

func main() {
	func(str string) {
		fmt.Println(str)
	}("你好，三酷猫")

	//声明变量functionA，并将匿名函数赋值给它
	functionA:=func(str string) {
		fmt.Println(str)
	}
	//通过functionA()调用匿名函数
	functionA("你好")
	functionA("三酷猫")

	//调用start()函数，传入匿名函数，输出当前时间
	start(func(){
		t:=time.Now()
		fmt.Println(t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	})

}

//命名函数start()，每隔一秒通过f()调用一次匿名函数
func start(f func()){
	for{
    //延迟1秒执行
		time.Sleep(1*time.Second)
		f()
	}
}
