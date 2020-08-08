package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello golang"
	err := ioutil.WriteFile("/Users/kvn/Desktop/go/学习/study30/test3.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
