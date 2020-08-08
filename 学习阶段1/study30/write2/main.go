package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("/Users/kvn/Desktop/go/学习/study30/test2.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	// writer.WriteString("你好golang") //将数据存入缓存
	for i := 0; i < 10; i++ {
		writer.WriteString("写入字符串" + strconv.Itoa(i) + "\r\n")
	}
	writer.Flush() //将缓存中的数据写入
}
