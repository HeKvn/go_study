package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string //私有属性不能被json包访问，要公有首字母大写
	Sno    string
}

func main() {

	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v\n", s1) //main.Student{ID:12, Gender:"男", Name:"李四", Sno:"s0001"}

	//结构体对象转换成json对象
	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr) //{"ID":12,"Gender":"男","Name":"李四","Sno":"s0001"}

	//将json字符串转换成结构体
	var str = `{"ID":10,"Gender":"男","Name":"王五","Sno":"s0002"}`
	var s2 Student
	err := json.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Println("失败")
	} else {
		fmt.Printf("----%#v\n", s2)
	}
	fmt.Println(s2.Name)

}
