package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
)

func main() {
	//请求地址和参数
	params:= url.Values{}
	apiUrl,err:=url.Parse("https://api.muxiaoguo.cn/api/tianqi")
	if err != nil {
		fmt.Println("解析请求地址出错，错误信息：", err)
		return
	}
	params.Set("type","1")
	params.Set("city","北京")
	//调用Encode()函数，为URL添加参数，并兼容中文字符
	apiUrl.RawQuery=params.Encode()
	//使用net包发起Get请求
	fmt.Println("请求地址：",apiUrl.String())
	resp, err := http.Get(apiUrl.String())
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
