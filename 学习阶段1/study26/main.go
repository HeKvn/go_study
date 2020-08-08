package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1()你好，golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2()你好，golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //计数器-1
}

func main() {
	/*
		为了防止主线程结束而协程还没完却直接退出的操作，使用sync.Waitgroup实现主线程等待协程执行完毕
		可以把wg.Add(1)写进for循环里
	*/
	wg.Add(2)  //协程计数器加2  写两个Add(1)也行的
	go test1() //表示开启一个协程
	// wg.Add(1)
	go test2()
	for i := 0; i < 10; i++ {
		fmt.Println("main()你好，golang-", i)
		time.Sleep(time.Millisecond * 50)
	}

	//查看电脑CPU个数
	cpuNum := runtime.NumCPU()
	fmt.Println("-------", cpuNum)
	//可以自己设置使用多个cpu
	// runtime.GOMAXPROCS(cpuNum-1)

	wg.Wait() //等待协程执行完毕
	fmt.Println("主线程退出")
}
