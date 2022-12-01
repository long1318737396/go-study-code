package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//构建自定义的Client变量
	client := &http.Client{}
	//定义HTTP请求的方式、地址和参数
	req,err:=http.NewRequest("GET","https://api.muxiaoguo.cn/api/lishijr",nil)
	//为本次请求添加Header信息
	req.Header.Add("headerKey1","headerValue1")
	//使用自定义的Client发起Get请求
	resp, err := client.Do(req)
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
