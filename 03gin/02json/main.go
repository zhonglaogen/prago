package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//参数传递
	router := gin.Default()

	router.GET("/json", func(context *gin.Context) {
		//方法1 map
		data := map[string]interface{}{
			"name": "zlx",
			"age": 20,
		}
		//方法2 结构体
		
		context.JSON(http.StatusOK,data)


	})

	type msg struct {
		//json包做反射序列化,所以写json
		Name string `json:"name"`
		Msg  string
		Age int
	}
	router.GET("/second", func(context *gin.Context) {
		//如果结构体小写字母没法传过去
		data := msg{
			"zlx",
			"hello" ,
			1,
		}
		context.JSON(http.StatusOK,data) //json 序列化使用go的序列化,通过反射
	})
	
	router.Run(":9090")
}
