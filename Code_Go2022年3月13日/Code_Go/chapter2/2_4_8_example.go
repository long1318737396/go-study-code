package main
import (
	"fmt"
)
/*三酷猫找数求和
*/
func main() {
	var num int=12345
	var numA=num/10000
	var numB=num%10000/1000
	var numC=num%1000/100
	var numD=num%100/10
	var numE=num%10
	fmt.Println(numA+numB+numC+numD+numE)
}
