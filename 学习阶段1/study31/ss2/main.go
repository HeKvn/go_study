package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFileName string, dstFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	defer sFile.Close()
	sFile2, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer sFile2.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	var tempSlice = make([]byte, 128)
	for {
		//读取数据
		n1, e1 := sFile.Read(tempSlice)
		if e1 == io.EOF {
			break
		}
		if e1 != nil {
			return e1
		}
		//写入数据
		_, e2 := sFile2.Write(tempSlice[:n1])
		if e2 != nil {
			return e2
		}
	}
	return nil
}

func main() {
	src := "/Users/kvn/Desktop/go/学习/study31/ps.txt"
	dst := "/Users/kvn/Desktop/go/学习/study31/ss2/ps.txt"
	err := copyFile(src, dst)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}

//生成目录（文件夹）的话用Mkdir、MkdirAll
//删除文件 Remove、RemoveAll  删除目录("./目录名") 删除文件("文件名")  一次删除多个文件RemoveAll("父目录")
//重命名 Rename("重命名文件的路径和文件名","重命名成功后的路径和文件名")
