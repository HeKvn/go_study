package main

import (
	"encoding/json"
	"fmt"
)

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

	c := Class{
		Title:   "001班",
		Student: make([]Student, 0),
	}

	for i := 1; i <= 10; i++ {
		s := Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Student = append(c.Student, s)
	}
	// fmt.Println(c) //{001班 [{1 男 stu_1} {2 男 stu_2} {3 男 stu_3} {4 男 stu_4} {5 男 stu_5} {6 男 stu_6} {7 男 stu_7} {8 男 stu_8} {9 男 stu_9} {10 男 stu_10}]}

	//将嵌套结构体转换成json字符串
	strByte, _ := json.Marshal(c)
	strJson := string(strByte)
	fmt.Println(strJson)
}
