package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//任意请求方法类型
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK,gin.H{
				"msg":"get",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{
				"msg": "post",
			})
		}
	})
	//错误请求返回地址
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"msg": "notfoud",
		})
	})
	//路由组的使用
	g1 := r.Group("/video")
	g1.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"hello",
		})
	})

	r.Run(":9090")
}
