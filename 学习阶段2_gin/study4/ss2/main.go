package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//设置html目录 否则无法访问静态html
	engine.LoadHTMLGlob("/Users/kvn/Desktop/go/学习/学习阶段2_gin/study4/html/*")
	//这里写相对路径报错pattern matches no files:

	engine.Static("/image", "/Users/kvn/Desktop/go/学习/学习阶段2_gin/study4/image")
	//都只能写绝对路径 哭了
	engine.GET("/hellohtml", func(context *gin.Context) {
		fullpath := "请求路径:" + context.FullPath()
		fmt.Println(fullpath)
		context.HTML(http.StatusOK, "index.html", gin.H{
			//html模版数据
			"fullpath": fullpath,
			"title":    "Gin学习",
		})
		//http.StatusOK相当于200
	})

	engine.Run()
	//run方法里可以指定端口，默认8080，可以修改成Run(":8090")
}
