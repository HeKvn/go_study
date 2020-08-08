package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//向intChan放入1-120000个数
func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

//向primeChan存放素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break //跳出这一级的循环，回到上一级循环
			}
		}
		if flag {
			primeChan <- num
		}
	}
	// for rang管道，要关闭管道，这里要关闭primeChan
	// close(primeChan) //但是如果一个channel关闭了，就没法给这个channel写入数据了
	//什么时候关闭primeChan?

	//给exitChan里面放入一条数据
	exitChan <- true
	wg.Done()
}

func printPrime(primeChan chan int) {
	// for v := range primeChan {
	// 	fmt.Println(v)
	// }
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000) //整个程序写完之后，没改50000还是1000之前，注释掉了打印方法，导致程序最后出错，原因：primeChan容量太小又只进没出，最后管道崩溃。打印方法有让primeChan数据流出的
	exitChan := make(chan bool, 16)    //标示primeChan close
	wg.Add(1)
	//存放数字的协程
	go putNum(intChan)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		//统计素数的协程
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	//打印素数的协程
	go printPrime(primeChan)

	//判断exitChan是否存满值 创建一个协程，使用匿名自执行函数来做
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		//总共就16个数据，如果前面的协程比较慢，这里会等待，然后取完之后就会退出循环
		//读取速度大于写入速度，读取的就会等待写入完成
		//关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("执行完毕")
	end := time.Now().Unix()
	fmt.Println(end-start, "秒")
}
