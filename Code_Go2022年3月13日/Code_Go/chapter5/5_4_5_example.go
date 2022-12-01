package main
import (
	"fmt"
	"time"
)
/*三酷猫的下载神器
 */
func main() {
	download(func(progress int) {
		fmt.Println("当前下载进度：",progress)
	})
}
func download(output func(progress int))  {
	value:=0
	for{
		output(value)
		value++
		if value>100{
			break
		}
		time.Sleep(1*time.Second)
	}
}
