package main

import "github.com/gin-gonic/gin"

func main() {
	//get 绑定
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {


		name := context.Query("query")

		context.JSON(200,name)
	})
	router.Run(":9000")
}
