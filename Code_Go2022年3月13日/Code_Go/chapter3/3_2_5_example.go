package main
import (
	"fmt"
)
/*三酷猫算平均温度
*/
func main() {
	arrMulti := [7][3]float64{
		{15,18},
		{16,20},
		{13,16},
		{14,20},
		{12,18},
		{7,20},
		{6,20}}
	arrMulti[0][2]=(arrMulti[0][0]+arrMulti[0][1])/2
	arrMulti[1][2]=(arrMulti[1][0]+arrMulti[1][1])/2
	arrMulti[2][2]=(arrMulti[2][0]+arrMulti[2][1])/2
	arrMulti[3][2]=(arrMulti[3][0]+arrMulti[3][1])/2
	arrMulti[4][2]=(arrMulti[4][0]+arrMulti[4][1])/2
	arrMulti[5][2]=(arrMulti[5][0]+arrMulti[5][1])/2
	arrMulti[6][2]=(arrMulti[6][0]+arrMulti[6][1])/2
	fmt.Println("周一最低温：",arrMulti[0][0],"；最高温：",arrMulti[0][1],"；当日平均温：",arrMulti[0][2])
	fmt.Println("周二最低温：",arrMulti[1][0],"；最高温：",arrMulti[1][1],"；当日平均温：",arrMulti[1][2])
	fmt.Println("周三最低温：",arrMulti[2][0],"；最高温：",arrMulti[2][1],"；当日平均温：",arrMulti[2][2])
	fmt.Println("周四最低温：",arrMulti[3][0],"；最高温：",arrMulti[3][1],"；当日平均温：",arrMulti[3][2])
	fmt.Println("周五最低温：",arrMulti[4][0],"；最高温：",arrMulti[4][1],"；当日平均温：",arrMulti[4][2])
	fmt.Println("周六最低温：",arrMulti[5][0],"；最高温：",arrMulti[5][1],"；当日平均温：",arrMulti[5][2])
	fmt.Println("周日最低温：",arrMulti[6][0],"；最高温：",arrMulti[6][1],"；当日平均温：",arrMulti[6][2])
}
