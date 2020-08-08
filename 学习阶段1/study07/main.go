package main

import (
	"fmt"
	"sort"
)

// func sumFn(x int, y int) int {
// 	sum := x + y
// 	return sum
// }

func sortStr(map1 map[string]string) (str string) {
	var slicekey []string
	for key, _ := range map1 {
		slicekey = append(slicekey, key)
	}
	sort.Strings(slicekey)

	// var str string
	for _, v := range slicekey {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

func main() {
	// sum1 := sumFn(12, 3)
	// fmt.Println(sum1)
	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
	}
	ss := sortStr(m1)
	fmt.Println(ss)
}
