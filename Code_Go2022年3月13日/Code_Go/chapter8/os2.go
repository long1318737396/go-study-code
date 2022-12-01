package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func main() {
	write()
	read()
}

//写文件
func write() {
	//打开文件，模式为创建并只写，权限为0777
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	//关闭这个文件操作（使用断言）
	defer file.Close()
	writer := bufio.NewWriter(file)
	//使用bufio写入30行字符串
	for i := 0; i < 30; i++ {
		writer.WriteString("你好，三酷猫！")
	}
	//将bufio缓冲区的内容一次写入磁盘
	writer.Flush()
}
//读文件
func read() {
	//打开test.txt
	file, err := os.Open("./test.txt")
	if err != nil {
		return
	}
	//关闭这个文件操作（使用断言）
	defer file.Close()
	reader := bufio.NewReader(file)
	var fileData []byte
	//使用bufio按行读取文件内容
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			//当err为io.EOF时，表明已经到达文件末尾，应结束读取
			if err == io.EOF {
				break
			}else{
				fmt.Println("读文件出错，错误信息：", err)
				return
			}
		}
		fileData=append(fileData, line...)
	}
	fmt.Println(string(fileData))
}
