package main

import (
	"fmt"
	"io/ioutil"
)

//自己编写一个函数，接收两个文件路径，scrFileName dstFileName
func copy(scrFileName string, dstFileName string) (err error) {
	byteStr, err1 := ioutil.ReadFile(scrFileName)
	if err1 != nil {
		return err1
	}
	err2 := ioutil.WriteFile(dstFileName, byteStr, 0666)
	if err2 != nil {
		return err2
	}
	fmt.Println("复制文件成功")
	return nil
}

func main() {
	src := "/Users/kvn/Desktop/go/学习/study31/ps.txt"
	dst := "/Users/kvn/Desktop/go/学习/study31/ss1/ps.txt"
	err := copy(src, dst)
	if err != nil {
		fmt.Println(err)
	}
}
