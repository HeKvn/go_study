package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(RequestInfos())
	/*
		可以把中间件这行代码写进请求里面 如：
		engine.GET("/query", RequestInfos(),func(context *gin.Context){})
		因为第二个是可变参数，所以没关系，这样写的中间件就不是全局的了
	*/

	engine.GET("/query", func(context *gin.Context) {
		fmt.Println("中间件next记号")
		context.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  context.FullPath(),
		})
	})

	engine.GET("/hello", func(context *gin.Context) {
		//todo
	})

	engine.Run()
}

//context.Next可以将中间件的代码一分为二，分两个顺序执行
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		fullpath := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求path：" + fullpath)
		fmt.Println("请求method：" + method)

		context.Next()

		fmt.Println("状态信息：", context.Writer.Status())
	}
}
