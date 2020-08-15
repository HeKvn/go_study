package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

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
