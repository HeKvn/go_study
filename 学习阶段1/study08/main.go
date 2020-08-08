package main

import "fmt"

type calcType func(int, int) int //表示定义一个calc的类型，其类型为一个————参数是两个int类型，返回值为int类型的[函数]

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, cb calcType) int {
	//此处cb作为参数 参向函数 （把函数作为另一个函数的参数）
	return cb(x, y) //函数作为返回值
}

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	var c calcType
	c = add
	fmt.Printf("c的类型：%T\n", c)
	s := calc(10, 5, add)
	fmt.Println(s)

	//还可以这样写 第三个参数直接匿名函数 匿名函数（没有函数名称）
	j := calc(3, 4, func(x, y int) int {
		return x * y
	})
	fmt.Println(j) //12

	var a = do("+")
	fmt.Println(a(12, 4)) //16

	//下面写一个匿名函数 让他自执行
	func(x, y int) {
		fmt.Println("test...", x+y)
	}(10, 1) //11

	//也可以让一个变量等于匿名函数
	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(2, 3))
}
