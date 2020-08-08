package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("/Users/kvn/Desktop/go/学习/study30/test.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//写入文件
	for i := 0; i < 10; i++ {
		file.WriteString("写入字符串" + strconv.Itoa(i) + "\r\n")
	}
}
