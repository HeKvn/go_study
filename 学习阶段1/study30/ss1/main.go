package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//bufio读取文件
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //表示一次读取一行
		if err == io.EOF {
			fileStr += str //防止有空的行引起bug
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)
}
