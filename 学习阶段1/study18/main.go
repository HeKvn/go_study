package main

import (
	"encoding/json"
	"fmt"
)

//嵌套的结构体字符串转换会结构体
type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title   string
	Student []Student
}

func main() {
	str := `{"Title":"001班","Student":[{"ID":1,"Gender":"男","Name":"stu_1"},{"ID":2,"Gender":"男","Name":"stu_2"},{"ID":3,"Gender":"男","Name":"stu_3"},{"ID":4,"Gender":"男","Name":"stu_4"},{"ID":5,"Gender":"男","Name":"stu_5"},{"ID":6,"Gender":"男","Name":"stu_6"},{"ID":7,"Gender":"男","Name":"stu_7"},{"ID":8,"Gender":"男","Name":"stu_8"},{"ID":9,"Gender":"男","Name":"stu_9"},{"ID":10,"Gender":"男","Name":"stu_10"}]}`
	var c = &Class{}
	err := json.Unmarshal([]byte(str), c)
	if err != nil {
		fmt.Println("失败")
	} else {
		fmt.Printf("%#v\n", c)
		fmt.Printf("%#v\n", c.Title)
	}

}
