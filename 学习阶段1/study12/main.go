package main

import (
	"fmt"
	"time"
)

//定时器
func main() {
	ticker := time.NewTicker(time.Second) //设置一个定时器每隔一秒获取一下当前时间
	// ticker.C  for range他只会返回一个参数value，写的时候注意
	n := 0
	for value := range ticker.C {
		n++
		fmt.Println(value)
		if n == 5 {
			ticker.Stop() //5次之后中止定时器
			break
		}
	}

	//此外 time.Sleep(Second)代表休眠一秒,也可以用休眠方法做定时器,比如:
	/*
		n:=0
		for{
			time.Sleep(Second)
			fmt.Println("我在执行定时任务")
			n++
			if n==5{
				break
			}
		}
	*/
}
