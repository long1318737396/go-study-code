package main
import "fmt"
func main(){
	fmt.Println("程序开始执行")
	defer fmt.Println("发生宕机后要运行的逻辑")
	panic("发生宕机！")
	fmt.Println("程序停止执行")

	fmt.Println("程序开始执行")	
	defer fmt.Println("发生宕机后要运行的逻辑")
	numA:=[]int{}
	fmt.Println(numA[5])
	fmt.Println("程序停止执行")

	fmt.Println("程序开始执行")
	defer func() {
		err:=recover()
		if err=="USB设备被拔出"{
			fmt.Println("USB设备被拔出，请重新插入")
		}
	}()
	panic("USB设备被拔出")

	fmt.Println("程序开始执行")
	defer func() {
		err:=recover()
			fmt.Println(err)
	}()
	numA:=[]int{}
	fmt.Println(numA[5])


}
