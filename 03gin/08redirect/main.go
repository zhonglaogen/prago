package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//	请求重定向
	r := gin.Default()
	r.GET("/index",redFunc)
	r.GET("/a", func(c *gin.Context) {
		//跳转到/b对应的路由处理函数
		c.Request.URL.Path = "/b" //把请求的uri修改
		r.HandleContext(c) //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"bbbbb",
		})
	})
	r.Run(":9090")
}
func redFunc(c *gin.Context) {
	//c.JSON(http.StatusOK,gin.H{
	//	"msg":"ok",
	//})
	c.Redirect(http.StatusMovedPermanently,"https://www.sougou.com")
}
