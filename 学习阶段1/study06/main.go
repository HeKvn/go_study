package main

import (
	"fmt"
	"strings"
)

//查看有多少个单词的程序

func main() {
	var str = "how do you do how do you do"
	var strList = strings.Split(str, " ")

	fmt.Println(strList)

	var strMap = make(map[string]int)
	for _, v := range strList {
		strMap[v]++
	}
	fmt.Println(strMap)
}
