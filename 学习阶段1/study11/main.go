package main

import (
	"fmt"
	"time"
)

func main() {
	// timeObj := time.Now()
	// var str = timeObj.Format("2006-01-02 15:04:05")   //这里写的是格式化的格式
	// fmt.Println(str) //2020-07-21 01:16:57

	//获取当前时间戳
	/*
		时间戳是自1970年1月1日（08:00:00）至当前时间的总毫秒数。被称为unix时间戳
	*/
	timeObj := time.Now()
	str := timeObj.Unix()
	fmt.Println(str) //1595265765

	//把时间戳转换成日期字符串
	str1 := time.Unix(str, 0)
	fmt.Println(str1) //2020-07-21 01:28:06 +0800 CST
	var str3 = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str3)

	//日期字符串转换成时间戳
	var str4 = "2020-07-21 01:28:06"
	var tmp = "2006-01-02 15:04:05"
	timeobj, _ := time.ParseInLocation(tmp, str4, time.Local)
	fmt.Println("日期转换：", timeobj.Unix())

	//时间操作函数
	var tobj1 = time.Now()
	fmt.Println(tobj1)
	tobj1 = tobj1.Add(time.Hour) //增加一小时
	fmt.Println(tobj1)
}
