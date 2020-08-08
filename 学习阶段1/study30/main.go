package main

import (
	"fmt"
	"io"
	"os"
)

//使用file.Open读取
func main() {
	// file, err := os.Open("./ps.txt")
	file, err := os.Open("./main.go")
	defer file.Close() //必须关闭文件流
	if err != nil {
		fmt.Println(err)
		return
	}
	//操作文件
	// fmt.Println(file) //&{0xc000054180}

	//读取文件里面的内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	// n, err := file.Read(tempSlice)
	// if err != nil {
	// 	fmt.Println("读取失败")
	// }
	// fmt.Printf("读取到了%v个字节\n", n)
	// fmt.Println(string(tempSlice))
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
			//err==io.EOF表示读取完毕
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		// fmt.Printf("读取到了%v个字节\n", n)
		strSlice = append(strSlice, tempSlice[:n]...) //[:n]防止并入的错乱，更有效的并入(注意写法)
	}
	fmt.Println(string(strSlice))
}
