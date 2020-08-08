package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Computer struct{}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
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
