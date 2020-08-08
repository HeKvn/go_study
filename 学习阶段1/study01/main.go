package main

import "fmt"

func main() {
	// // var a int = 3
	// // fmt.Print(a, "\n")
	// // 输入换行符可以解决 zsh打印后的百分号，因为zsh如果输出末尾没有新的空行，会打出一个 % 作为提示。
	// var str = "this"
	// fmt.Printf("%v - 值%c\n", str[2], str[2])
	// var arr1 = []int{1, 2, 3, 4}
	// fmt.Println(arr1)
	// s := []int{1, 2, 3, 4, 5}
	// a := s[1:]
	// fmt.Printf("a长度%d 容量%d s长度%d s容量%d", len(a), cap(a), len(s), cap(s))
	var numSlice = []int{9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	fmt.Println(numSlice)
}
