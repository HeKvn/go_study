package main

import "fmt"

func main() {
	//map也是引用数据类型
	// make创建map类型的数据
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)

	//2.
	var test = map[string]string{
		"username": "张三",
		"age":      "23",
		"sex":      "男",
	}
	fmt.Println(test["age"])

	//3.
	kk := map[string]int{
		"kk1": 1,
		"kk2": 2,
		"kk3": 3,
	}
	//修改数据
	kk["kk1"] = 4
	fmt.Println(kk["kk1"])
	//查找
	v, ok := kk["kk1"]
	fmt.Println(v, ok)

	//for range 循环遍历映射
	for k, v := range kk {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	//删除
	dd := map[string]int{
		"dd1": 1,
		"dd2": 2,
		"dd3": 3,
	}
	delete(dd, "dd1")
	fmt.Println(dd)
}
