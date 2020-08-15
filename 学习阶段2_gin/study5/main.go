package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	routerGroup := engine.Group("/user")
	routerGroup.POST("/register", registerHandle)

	routerGroup.POST("/login", func(context *gin.Context) {
		fullpath := "用户登陆" + context.FullPath()
		fmt.Println(fullpath)
		context.Writer.WriteString(fullpath)
	})

	routerGroup.DELETE("/:id", func(context *gin.Context) {
		fullpath := "用户删除" + context.FullPath()
		userID := context.Param("id")
		fmt.Println(fullpath + " " + userID)
		context.Writer.WriteString(fullpath + " " + userID)
	})

	engine.Run()
}
func registerHandle(context *gin.Context) {
	fullpath := "用户注册" + context.FullPath()
	fmt.Println(fullpath)
	context.Writer.WriteString(fullpath)
}
