package main

import "fmt"

func main() {
	//定义一个map类型的切片  此处把map[string]string看成一个整体， 就相当于 int string
	var userinfo = make([]map[string]string, 3, 3)
	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["sex"] = "男"
	}

	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "23"
		userinfo[1]["sex"] = "男"
	}

	fmt.Println(userinfo)

	//value 可以是切片  此处把[]string看成一个整体
	var userinfo1 = make(map[string][]string)
	userinfo1["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}
	userinfo1["work"] = []string{
		"java",
		"go",
		"css",
	}
	fmt.Println(userinfo1)
	for k, v := range userinfo1 {
		for _, value := range v {
			fmt.Println(k, value)
		}
	}
}
