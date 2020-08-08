package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//ioutil读取
	byteStr, err := ioutil.ReadFile("/Users/kvn/Desktop/kvn/Markdowm/wx_cloud_init.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}
