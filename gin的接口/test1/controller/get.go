package controller

import (
	"fmt"

	"../model"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	i := model.Person{}
	err, is := i.List()
	if err != nil {
		fmt.Println("----", err)
		c.JSON(200, nil)
		return
	}
	c.JSON(200, is)
}
