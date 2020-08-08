package main

import (
	"fmt"
	"sort"
)

func main() {
	var list = make(map[int]int)
	list[1] = 4
	list[2] = 9
	list[5] = 44
	list[9] = 56
	list[4] = 30

	var keylist []int
	for key, _ := range list {
		keylist = append(keylist, key)
	}

	// fmt.Println(keylist)

	//让key进行升序排序
	sort.Ints(keylist)
	for _, val := range keylist {
		fmt.Printf("key:%v--value:%v\n", val, list[val])
	}

}
