package main

import "fmt"

//接口的详细使用
type Usber interface {
	start()
	stop()
}

type Computer struct{}

//在这里，判断传入的接口是手机还是什么别的，如果是手机，就调用start方法
func (c Computer) work(usb Usber) {
	if _, ok := usb.(Phone); ok {
		usb.start()
	} else {
		usb.stop()
	}
}

type Phone struct {
	Name string
}

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
	var computer = Computer{}
	var phone = Phone{
		Name: "小米手机",
	}
	computer.work(phone)
}
