package main

import "fmt"

type A interface{} //空接口，表示没有任何约束 任意的类型都可以实现空接口

//golang中空接口可以表示任何类型
func show(s interface{}) {
	fmt.Println(s)
}

func main() {
	var a A
	str := "你好，golang"
	a = str //让字符串实现A这个 接口
	fmt.Println(a)

	//golang中空接口可以表示任何类型
	var b interface{}
	b = 20
	fmt.Println(b)
	b = "nihao"
	fmt.Println(b)
	b = true
	fmt.Println(b)

	show("++++")
	slice := []int{1, 2, 3, 4, 5}
	show(slice)

	//ps:空接口就是 Object
	var m1 = make(map[string]interface{})
	m1["name"] = "张三"
	m1["age"] = 20
	fmt.Println(m1)
}
