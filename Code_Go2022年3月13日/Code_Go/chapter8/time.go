package main

import (
	"time"
	"fmt"
)

func main() {
	now:=time.Now()
	fmt.Println(now)
	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	hour:=now.Hour()
	minute:=now.Minute()
	second:=now.Second()
	fmt.Printf("现在时刻：%d年%02d月%02d日 %02d时%02d秒%02d分\n", year, month, day, hour, minute, second)
	fmt.Println("Unix时间戳（秒）：",now.Unix())
	fmt.Println("Unix时间戳（毫秒）：",now.UnixMilli())
	fmt.Println("Unix时间戳（纳秒）：",now.UnixNano())

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 3:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04:05 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05"))

	oneHourLater:=now.Add(time.Hour)
	fmt.Println(oneHourLater)
	duration:=now.Sub(oneHourLater)
	fmt.Println(duration)

	fmt.Println(now.Equal(oneHourLater))
	fmt.Println(now.Before(oneHourLater))
	fmt.Println(now.After(oneHourLater))

	//设置时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("设置时区出现错误，错误信息：",err)
		return
	}
	//解析时间
	timeVal, err := time.ParseInLocation("2006/01/02 15:04:05", "2021/11/10 16:37:20", location)
	if err != nil {
		fmt.Println("解析时间出现错误，错误信息：",err)
		return
	}
	fmt.Println(timeVal)

	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println("你好，三酷猫！",i)
	}


}
