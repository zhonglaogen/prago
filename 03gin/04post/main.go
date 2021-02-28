package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//post绑定
	router := gin.Default()

	router.LoadHTMLFiles("./login.html","./index.html")
	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",nil)
	})
	router.POST("login", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		if username=="hello" && password=="123" {
			context.JSON(200,"ok")
		}else {
			context.HTML(200,"index.html",gin.H{
				"Name": username,
				"password": password,
			})
		}

	})

	router.Run(":9000")
}
