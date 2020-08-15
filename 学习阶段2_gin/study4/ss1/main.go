package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellojson", func(context *gin.Context) {
		fullpath := "请求路径:" + context.FullPath()
		fmt.Println(fullpath)
		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "ok",
			"data":    fullpath,
		})
	})

	engine.GET("/jsonstruct", func(context *gin.Context) {
		fullpath := "请求路径:" + context.FullPath()
		fmt.Println(fullpath)

		resp := Respone{Code: 1, Message: "okk", Data: fullpath}
		context.JSON(200, &resp)
		//传入地址，传入真身
	})

	engine.Run()
}

type Respone struct {
	Code    int
	Message string
	Data    interface{}
}
