package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//param绑定
	router := gin.Default()
	//防止冲突要加上前缀
	router.GET("/user/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK,gin.H{
			"name":name,
			"age": age,
		})

	})
	router.GET("/blog/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK,gin.H{
			"name":name,
			"age": age,
		})

	})
	router.Run(":9090")
}
