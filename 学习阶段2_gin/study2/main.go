package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//处理一个http://localhost:8080/hello?name=davie
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)
		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("Hello," + name))
	})

	//post
	//http://localhost:8080/login
	engine.Handle("POST", "login", func(contex *gin.Context) {
		fmt.Println(contex.FullPath())

		username := contex.PostForm("username")
		password := contex.PostForm("password")
		fmt.Println(username)
		fmt.Println(password)
		contex.Writer.Write([]byte(username + "登陆"))
	})

	engine.Run()
}
