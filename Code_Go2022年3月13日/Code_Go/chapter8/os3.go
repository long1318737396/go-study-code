package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	write()
	read()
}

//写文件
func write() {
	err := ioutil.WriteFile("./test.txt", []byte("你好，三酷猫！"), os.ModePerm)
	if err != nil {
		fmt.Println("写文件操作出现错误，异常信息：", err)
		return
	}
}

//读文件
func read() {
	fileData, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("读文件操作出现错误，异常信息：", err)
		return
	}
	fmt.Println(string(fileData))
}
