package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//请求地址
	url := "https://res.abeim.cn/api-baidu_keyword"
	//内容类型
	contentType := "application/x-www-form-urlencoded"
	//表单数据
	data:="wd=golang"
	//使用net包发起Post请求
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("HTTP请求失败，错误信息：", err)
		return
	}
	//使用断言关闭网络请求
	defer resp.Body.Close()
	//使用ioutil工具包获取服务端响应数据
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络响应失败，错误信息：", err)
		return
	}
	fmt.Println(string(b))
}
