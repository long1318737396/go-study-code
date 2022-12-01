package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func genQrCode(content string) {
	resp, err :=
		http.Get("https://api.muxiaoguo.cn/api/Qrcode?frame=1&e=L&text=" + content + "&size=500")
	if err != nil {
		fmt.Println("HTTP请求失败，错误信息：", err)
		return
	}
	//使用断言关闭网络请求
	defer resp.Body.Close()
	//使用bufio读取网络响应
	reader := bufio.NewReaderSize(resp.Body, 32*1024)
	//创建目标文件
	file, err := os.Create("./qrcode.png")
	if err != nil {
		panic(err)
	}
	//向目标文件写入数据
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Printf("图片文件已保存，文件大小： %d", written)
}

func main() {
	genQrCode("你好，三酷猫")
}
