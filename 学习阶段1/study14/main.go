package main

import "fmt"

//定义结构体
type person struct {
	name string
	age  int
	sex  string
}

//结构体方法
func (p person) printInfo() {
	fmt.Printf("姓名：%v  年龄：%v\n", p.name, p.age)
}

//这里要注意，要修改结构体里的数据，接收者类型是指针类型，即*person，否则会修改不成功
//记住，修改结构体数据都要用指针类型的接收者类型
func (p *person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

//除此之外，在golang中还可以自定义类型，然后给自定义的类型添加方法，就像结构体一样
//注意：非本地类型不能自定义方法，我们只能给自己写的包里的自定义类型添加方法
type Myint int

func (m Myint) Printinfo() {
	fmt.Println("我是自定义类型里的自定义方法")
}

func main() {
	var p1 person //实例化person结构体
	p1.name = "张三"
	p1.age = 20
	// p1.sex = "男"
	// fmt.Printf("值：%v--类型：%T\n", p1, p1)
	// fmt.Printf("值：%#v--类型：%T\n", p1, p1)
	p1.printInfo()

	var a Myint = 20

	a.Printinfo()
}
