package main

import "fmt"

//x.(type)来判断类型 这个语句只能用在switch里面
func MyPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("其他类型")

	}
}

//类型断言
func main() {
	var a interface{}
	a = "你好golang"
	_, ok := a.(string)
	if ok {
		fmt.Println("a就是strint类型，值是：", a)
	} else {
		fmt.Println("断言失败   ")
	}
	MyPrint(23)
}
