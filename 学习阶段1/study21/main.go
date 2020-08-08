package main

import "fmt"

//接口是一个约定，一个规范
type Usber interface {
	start()
	stop()
}

//如果接口里面有方法的话，必须通过结构体或自定义类型实现这个接口

type Phone struct {
	Name string
}

//手机要实现 usb接口必须得实现usb接口里的所有方法
//开始定义结构体方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关闭")
}

//允许有接口以为的结构体方法
func (p Phone) run() {
	fmt.Println(p.Name, "运行程序")
}

func main() {
	p := Phone{
		Name: "华为手机",
	}
	// p.start()
	var p1 Usber
	p1 = p
	p1.start()
	//但没办法通过接口运行规范结构体之外的方法，如：p1.run()
}
