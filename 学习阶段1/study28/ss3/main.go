package main

import (
	"fmt"
	"time"
)

//多路复用

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//使用select来获取channel里的数据不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%v\n", v)
		default:
			fmt.Println("数据获取完毕")
			return //注意退出
		}
	}
}
