package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(Cors())

	//http://localhost:8080/hello?name=davie&classes=软件工程
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var student Student
		//将请求回来的数据绑定到student结构体里，一一映射
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		//这样返回数据
		context.JSON(200, map[string]interface{}{
			"name":    student.Name,
			"jiaoshi": student.Classes,
		})
		//打印一下
		// context.Writer.Write([]byte("hello," + student.Name))
	})

	engine.Run()
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

//解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			// 设置请求头，这里是重点！！。然后前端可以传递一个请求头mytoken，比如登陆的时候有用的
			c.Header("Access-Control-Allow-Headers", "mytoken")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
