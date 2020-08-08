package main

import (
	"fmt"
	"sort"
)

func main() {

	//升序排序
	// intList := []int{2, 5, 6, 1, 3, 8, 7, 9, 4}
	// sort.Ints(intList) //sort包可以直接排序
	// fmt.Println(intList)
	// stringList := []string{"a", "c", "e", "d", "b"}
	// sort.Strings(stringList)
	// fmt.Println(stringList)

	//降序排序
	intList := []int{2, 5, 6, 1, 3, 8, 7, 9, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)
}
