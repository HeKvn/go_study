package main

import "fmt"

func test() string {
	return "hello"
}

func main() {
	fmt.Println("hello world")
	var a int = 10
	var b int = 3
	var c int = 2
	fmt.Printf("%v %v %v", a, b, c)
	// b := 2
	// fmt.Print(b)
	test()
}
