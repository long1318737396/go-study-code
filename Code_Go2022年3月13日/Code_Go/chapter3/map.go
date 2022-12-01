package main
import (
	"fmt"
	"reflect"
)
/*Map类型
*/
func main() {
	//声明map型变量mapValueA
	var mapValueA map[string]string
	fmt.Println("mapValueA的类型为：",  reflect.TypeOf(mapValueA))
	fmt.Println("mapValueA==nil?", mapValueA==nil)
	fmt.Println("mapValueA的长度为：", len(mapValueA))

	//使用make()函数初始化集合
	mapValueB := make(map[string]string)
	mapValueB["key1"] = "value1"
	mapValueB["key2"] = "value2"
	fmt.Println("mapValueB的值为：", mapValueB)
	fmt.Println("mapValueB的长度为：", len(mapValueB))

	//直接给定具体值初始化集合
	var mapValueC = map[string]string{"key1": "value1", "key2": "value2"}
	fmt.Println("mapValueC的值为：", mapValueC)
	fmt.Println("mapValueC的长度为：", len(mapValueC))
	fmt.Println("获取键key1的值", mapValueC["key1"])
	fmt.Println("获取键key3的值", mapValueC["key3"])
	fmt.Println(mapValueC["key3"]=="")
	//获取键为key3的值并判断是否存在这样的值
	key3Value, key3Exists:=mapValueC["key3"]
	fmt.Println(key3Exists,key3Value)

	//直接给定具体值初始化集合
	var mapValueD = map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3"}
	fmt.Println("mapValueD的值为：", mapValueD)
	//向mapValueD中添加元素
	mapValueD["key4"]="value4"
	mapValueD["key5"]="value5"
	fmt.Println("mapValueD的值为：", mapValueD)
	//删除键为key5的键值对
	delete(mapValueD,"key5")
	fmt.Println("mapValueD的值为：", mapValueD)
	//修改键为key1的值
	mapValueD["key1"]="VALUE1"
	fmt.Println("mapValueD的值为：", mapValueD)
}
