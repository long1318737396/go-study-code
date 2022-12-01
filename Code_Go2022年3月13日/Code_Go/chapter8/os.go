package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//创建一个文件，目录为./test.txt（./表示当前目录）
	file, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println("创建文件失败，错误信息：", err)
		return
	}

	count, err := file.WriteString("你好，三酷猫")
	fmt.Println("总计写入：", count)
	if err != nil {
		fmt.Println("写入文件异常，错误信息：", err)
		return
	}

	//关闭文件操作
	file.Close()

	//打开一个文件，目录为./test.txt（./表示当前目录）
	file, err = os.Open("./test.txt")
	if err != nil {
		fmt.Println("打开文件失败，错误信息：", err)
		return
	}
	//关闭文件操作
	defer file.Close()

	//声明oneTimeBuf，用作循环中读取文件时的缓冲
	var oneTimeBuf [512]byte
	//声明fileData，用作保存文件完整数据
	var fileData []byte
	for {
		count, err := file.Read(oneTimeBuf[:])
		if err != nil {
			//当err为io.EOF时，表明已经到达文件末尾，应结束读取
			if err == io.EOF {
				fmt.Println("读文件结束。")
				break
			} else {
				fmt.Println("读文件出错，错误信息：", err)
				return
			}
		}
		fileData = append(fileData, oneTimeBuf[:count]...)
	}
	fmt.Println(string(fileData))

	err = os.Remove("./test.txt")
	if err != nil {
		fmt.Println("删除文件失败，错误信息：", err)
		return
	}

	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("文件存在")
			fmt.Println("这个文件是目录吗？", fileInfo.IsDir())
		} else {
			fmt.Println("文件不存在")
		}
	} else {
		fmt.Println("文件存在")
		fmt.Println("这个文件是目录吗？", fileInfo.IsDir())
	}

}
