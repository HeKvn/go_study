package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//参数写在body里

func main() {
	engine := gin.Default()
	engine.Use(Cors())

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		err := context.ShouldBind(&register)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Password)
		context.Writer.Write([]byte(register.UserName + " Register"))
	})

	engine.Run()
}

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
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
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
