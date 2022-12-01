package main
import (
	"fmt"
	"reflect"
)
/*切片类型
*/
func main() {
	//声明并使用make()函数初始化arrZeroInt切片变量
	var arrZeroInt=make([]int32, 4, 6)
	fmt.Println("arrZeroInt=", arrZeroInt)
	//对切片变量使用Go语言内置的len()函数来获取切片长度
	fmt.Println("arrZeroInt的长度为", len(arrZeroInt))
	//对切片变量使用Go语言内置的cap()函数来获取切片容量
	fmt.Println("arrZeroInt的容量为", cap(arrZeroInt))

	//声明arrInt数组变量，并初始化
	var arrInt=[6]int{11, 15, 25, 23, 19, 78}
	fmt.Println("arrInt=", arrInt)
	fmt.Println("arrInt的长度为", len(arrInt))
	//声明并节选arrInt数组用作初始化sliceInt切片
	var sliceInt=arrInt[1:5]

	// var sliceInt=arrInt[1:]
	// var sliceInt=arrInt[:5]

	fmt.Println("sliceInt=", sliceInt)
	//对切片变量使用Go语言内置的len()函数来获取切片长度
	fmt.Println("sliceInt的长度为", len(sliceInt))
	//对切片变量使用Go语言内置的cap()函数来获取切片容量
	fmt.Println("sliceInt的容量为", cap(sliceInt))

	var sliceInt2=[]int{1,2,3,4,5}
	fmt.Println("sliceInt2=", sliceInt2)
	//对切片变量使用Go语言内置的len()函数来获取切片长度
	fmt.Println("sliceInt2的长度为", len(sliceInt2))
	//对切片变量使用Go语言内置的cap()函数来获取切片容量
	fmt.Println("sliceInt2的容量为", cap(sliceInt2))

	var sliceValue []int
	fmt.Println(sliceValue==nil)
	fmt.Println(len(sliceValue))

	//函数体内使用:=赋值运算符声明int型数组sliceInt3
	//声明长度为6，数据项类型为int的数组sliceInt3并初始化
	sliceInt3 := []int{11, 15, 25, 23, 19, 78}
	fmt.Println("sliceInt3=", sliceInt3)
	fmt.Println("sliceInt3的长度为", len(sliceInt3))
	//获取sliceInt的第一个元素
	fmt.Println("sliceInt3的第一个元素为", sliceInt3[0])
	//可以通过Go语言标准库的reflect包中的TypeOf函数查看数组sliceInt3的数据类型
	fmt.Println("sliceInt3的数据类型为", reflect.TypeOf(sliceInt3))

	//声明一个切片变量sliceInt4
	var sliceInt4 []int
	fmt.Println("sliceInt4=", sliceInt4)
	fmt.Println("sliceInt4的长度为", len(sliceInt4))
	fmt.Println("sliceInt4的容量为", cap(sliceInt4))
	fmt.Printf("sliceInt4的地址为 %p \n", sliceInt4)
	//向切片中追加1个元素
	sliceInt4=append(sliceInt4,1)
	fmt.Println("sliceInt4=", sliceInt4)
	fmt.Println("sliceInt4的长度为", len(sliceInt4))
	fmt.Println("sliceInt4的容量为", cap(sliceInt4))
	fmt.Printf("sliceInt4的地址为 %p \n", sliceInt4)
	//向切片中追加多个元素
	sliceInt4=append(sliceInt4,2,3,4)
	fmt.Println("sliceInt4=", sliceInt4)
	fmt.Println("sliceInt4的长度为", len(sliceInt4))
	fmt.Println("sliceInt4的容量为", cap(sliceInt4))
	fmt.Printf("sliceInt4的地址为 %p \n", sliceInt4)
	//向切片中追加另一个切片
	sliceInt4=append(sliceInt4,[]int{5,6,7}...)
	fmt.Println("sliceInt4=", sliceInt4)
	fmt.Println("sliceInt4的长度为", len(sliceInt4))
	fmt.Println("sliceInt4的容量为", cap(sliceInt4))
	fmt.Printf("sliceInt4的地址为 %p \n", sliceInt4)

	//声明一个切片变量sliceInt5
	var sliceInt5 []int
	fmt.Println("sliceInt5=", sliceInt5)
	fmt.Println("sliceInt5的长度为", len(sliceInt5))
	fmt.Println("sliceInt5的容量为", cap(sliceInt5))
	fmt.Printf("sliceInt5的地址为 %p \n", sliceInt5)
	//向切片开头追加多个元素
	sliceInt5=append([]int{-2,-1,0},sliceInt5...)
	fmt.Println("sliceInt5=", sliceInt5)
	fmt.Println("sliceInt5的长度为", len(sliceInt5))
	fmt.Println("sliceInt5的容量为", cap(sliceInt5))
	fmt.Printf("sliceInt5的地址为 %p \n", sliceInt5)

	//声明一个切片变量sliceInt6
	var sliceInt6 []int
	fmt.Println("sliceInt6=", sliceInt6)
	fmt.Println("sliceInt6的长度为", len(sliceInt6))
	fmt.Println("sliceInt6的容量为", cap(sliceInt6))
	fmt.Printf("sliceInt6的地址为 %p \n", sliceInt6)
	//嵌套使用append()函数
	sliceInt6=append(sliceInt6,append([]int{-2,-1,0},append(sliceInt6,1)...)...)
	fmt.Println("sliceInt6=", sliceInt6)
	fmt.Println("sliceInt6的长度为", len(sliceInt6))
	fmt.Println("sliceInt6的容量为", cap(sliceInt6))
	fmt.Printf("sliceInt6的地址为 %p \n", sliceInt6)

	//声明一个切片变量sliceInt7
	var sliceInt7 []int=[]int{0,1,4}
	fmt.Println("sliceInt7=", sliceInt7)
	fmt.Println("sliceInt7的长度为", len(sliceInt7))
	fmt.Println("sliceInt7的容量为", cap(sliceInt7))
	fmt.Printf("sliceInt7的地址为 %p \n", sliceInt7)
	//在切片的第二个元素开始插入元素
	sliceInt7=append(sliceInt7[:2],append([]int{2,3},sliceInt7[2:]...)...)
	fmt.Println("sliceInt7=", sliceInt7)
	fmt.Println("sliceInt7的长度为", len(sliceInt7))
	fmt.Println("sliceInt7的容量为", cap(sliceInt7))
	fmt.Printf("sliceInt7的地址为 %p \n", sliceInt7)

	//声明切片变量sliceA和slice2B并赋初始值
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{5, 4, 3}
	// 只会复制sliceA的前3个元素到sliceB中
	copy(sliceB, sliceA)
	fmt.Println("sliceB的值为：",sliceB)

	// 只会复制sliceB的3个元素到sliceA的前3个位置
	// copy(sliceA, sliceB)
	// fmt.Println("sliceA的值为：",sliceA)


}
