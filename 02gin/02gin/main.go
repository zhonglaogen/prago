package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main ()  {
//	创建一个默认的路由引擎
	r := gin.Default()
//	get请求方式,/hello为请求的路径
	r.GET("/hello",sayhello)
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"method":"get",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"method":"post",
		})
	})
	//r.PUT()
	//r.DELETE()

//	启动服务,什么都不写默认是8080
	r.Run(":9090");



}
func sayhello(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"hello world",
	})
}

