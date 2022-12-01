package main
import (
	"fmt"
)

func main() {
	var sum int = 0
	var i int
	for i=1;i<=100;i++{
		sum+=i
		if i==10{
			break
		}
	}
	fmt.Println(sum)

	var j int
	var k int
	out:
	for j=1;j<=3;j++{
		for k=10;k<=15;k++{
			fmt.Println("j=",j,",k=",k)
			if k==12{
				break out
			}
		}
	}

	var i2 int
	for i2=0;i2<=5;i2++{
		if i2>3&&i2<5{
			continue
		}
	fmt.Println("i2=",i2)
	}

	var i3 int
	var j3 int
	for i3 = 0; i3 <= 5; i3++ {
		for j3 = 0; j3 <= 5; j3++ {
			if i3 == 1 {
				goto stopLoop
			}
			fmt.Println("i3=", i3, ",j3=", j3)
		}
	}
stopLoop:
	fmt.Println("i3=1，结束循环")

	
}