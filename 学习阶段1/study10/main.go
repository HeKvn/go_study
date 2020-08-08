package main

import "fmt"

func f1() {
	fmt.Println("开始")
	defer func() {
		fmt.Println("aaa")
	}()
	fmt.Println("结束")
}

func main() {
	f1()
}
