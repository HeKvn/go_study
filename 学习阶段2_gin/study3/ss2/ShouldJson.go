package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		err := context.BindJSON(&person)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println("姓名:" + person.Name)
		context.Writer.Write([]byte("添加记录：" + person.Name))
	})

	engine.Run()

}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}
