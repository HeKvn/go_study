package main

import (
	"encoding/json"
	"fmt"
)

//转换完之后让首字母小写  结构体标签
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr) //{"id":12,"gender":"男","name":"李四","sno":"s0001"}

}
