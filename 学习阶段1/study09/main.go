package main

import "fmt"

/*
	全局变量：常驻内存，污染全局
	局部变量：不常驻内存，不污染全局
*/
//闭包
//常驻内存，又不污染全局
func adder1() func() int {
	var i = 10
	//这里的i相对于下面这个匿名函数来说，是全局变量，闭包，永远的神
	return func() int {
		return i + 1
	}
}

func adder2() func(int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	// var fn = adder1() //表示执行方法  其要返回一个函数，即fn其实是adder返回的一个函数，理解这一点，对下面的程序有帮助
	// fmt.Println(fn()) //11
	// fmt.Println(fn()) //11
	// fmt.Println(fn()) //11

	var fn2 = adder2()
	/*
		参上解释，fn2其实是一个adder2的返回函数，就是func(y int) int{}，
		要理解，当其用=连接时，就是执行了adder2这个函数，其将得到的是adder2的返回值，而这个返回值就是函数
	*/
	fmt.Println(fn2(1))
	fmt.Println(fn2(1))
	fmt.Println(fn2(1))
}
