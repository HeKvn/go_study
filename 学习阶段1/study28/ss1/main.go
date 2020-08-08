package main

import "fmt"

func main() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1)
	//使用 for range循环遍历管道的话要关闭管道，否则遍历完最后一个会报错。。而通过for循环来遍历的话可以不关闭管道

	//fo range循环遍历管道的值 注意，管道没有key
	for val := range ch1 {
		fmt.Println(val)
	}
}
