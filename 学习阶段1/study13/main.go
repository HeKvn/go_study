package main

import "fmt"

func main() {

	//在计算机底层 a这个变量其实对应了一个内存地址
	// var a = 10
	// fmt.Printf("a的值：%v a的类型%T a的地址%p\n", a, a, &a) //a的值：10 a的类型int a的地址0xc0000b2008

	var a = 10
	var b = &a //b的值0xc000014080 b的类型*int
	fmt.Printf("a的值%v a的类型%T a的地址%p\n", a, a, &a)
	fmt.Printf("b的值%v b的类型%T b的地址%p\n", b, b, &b)
	//golang里面变量都有一个对应的内存地址

	//通过指针b打印出a的值
	// *b //表示取出b这个变量对应的内存地址的值
	fmt.Println("---", *b) //--- 10

	//改变b变量对应的内存地址的值
	*b = 30
	fmt.Println(a)
}
