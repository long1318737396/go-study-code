package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	log.Println("打印日志，结尾自动换行")
	log.Print("打印日志，第二行\n")
	log.Printf("打印日志，第%s行","三")
	
	log.SetPrefix("log包测试")
	log.Println("日志前缀测试")

	if log.Flags()&log.Ldate==1{
		log.Print("log.Ldate已打开")
	}
	log.SetFlags(log.Llongfile|log.Ltime)
	if log.Flags()&log.Ldate==0{
		log.Print("log.Ldate已关闭")
	}
	
	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开日志文件出错，错误信息：", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("程序启动了")
	log.Fatalln("程序意外地崩溃了")
	
	customLogger := log.New(os.Stdout, "customLogger", log.Lshortfile|log.LstdFlags)
	customLogger.Println("使用自定义Logger输出日志")

}
