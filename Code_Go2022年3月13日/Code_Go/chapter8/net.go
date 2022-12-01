package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//使用net包发起Get请求
	resp, err := http.Get("https://api.muxiaoguo.cn/api/lishijr")
	if err != nil {
		fmt.Println("HTTP请求失败，错误信息：", err)
		return
	}
	//使用断言关闭网络请求
	defer resp.Body.Close()
	//使用ioutil工具包获取服务端响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络响应失败，错误信息：", err)
		return
	}
	fmt.Print(string(body))
}
