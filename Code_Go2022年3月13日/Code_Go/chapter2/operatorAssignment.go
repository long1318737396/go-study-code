package main
import (
	"fmt"
)
/*赋值运算符的使用
*/
func main() {
	var numOne int=20                  //声明numOne变量
	numOne+=20                       //numOne=numOne+20
	fmt.Println(numOne)                 //输出numOne的值
	numOne-=10                       //numOne=numOne-10
	fmt.Println(numOne)                 //输出numOne的值
	numOne*=100                      //numOne=numOne*100
	fmt.Println(numOne)                 //输出numOne的值
	numOne/=20                       //numOne=numOne/20
	fmt.Println(numOne)                 //输出numOne的值
	numOne%=4                       //numOne=numOne&4
	fmt.Println(numOne)                 //输出numOne的值
	numOne<<=2                       //numOne=numOne<<2
	fmt.Println(numOne)                 //输出numOne的值
	numOne>>=3                       //numOne=numOne>>3
	fmt.Println(numOne)                 //输出numOne的值
	numOne&=0                        //numOne=numOne&0
	fmt.Println(numOne)                 //输出numOne的值
	numOne^=1                        //numOne=numOne^1
	fmt.Println(numOne)                 //输出numOne的值
	numOne|=0                        //numOne=numOne|0
	fmt.Println(numOne)                 //输出numOne的值
}
