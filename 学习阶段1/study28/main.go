package main

import "fmt"

func main() {
	//创建管道
	ch := make(chan int, 3)
	//给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32
	//获取管道里面的内容
	a := <-ch
	fmt.Println(a) //10
	<-ch
	c := <-ch
	fmt.Println(c) //32
	/*
		管道遵循先入先出
	*/
	//再存值
	ch <- 56
	//打印管道的长度和容量
	fmt.Printf("值：%v  容量：%v 长度%v\n", ch, cap(ch), len(ch)) //值：0xc000022080  容量：3 长度1
	//管道的类型(引用数据类型)
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64
	ch2 := ch1
	ch2 <- 25
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d)

	//管道阻塞
	// ch6 := make(chan int, 1)
	// ch6 <- 34
	// ch6 <- 43
	// //all goroutines are asleep - deadlock!
	ch7 := make(chan string, 2)
	ch7 <- "数据1"
	ch7 <- "数据2"
	m1 := <-ch7
	m2 := <-ch7
	fmt.Println(m1, m2)
	//此外，在没有使用协程的情况下，管道的数据已经取完之后，再取就会报错的
}
