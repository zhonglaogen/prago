package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("============")
	value, exists := c.Get("name") //获得上下文值
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"msg": value,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	}

}
func test(c *gin.Context) {
	fmt.Println("*****************")
	// go funcxx()  要用c.Copy(),防止数据混乱

	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

//中间件
func midHandler(c *gin.Context) {
	fmt.Println("mid ^^^^^")

	//	计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n",cost)

}

//中间件
func m2(c *gin.Context) {
	fmt.Println("inm222222222222")


	c.Next() //调用后续的处理函数

	fmt.Println("outm222222222222")

}
//闭包
func autoMidlleware(doCheck bool)gin.HandlerFunc  {
	//连接数据库,其他准备工作中
	return func(c *gin.Context) {
		if doCheck {
			//xxxxxxx
		//	存放具体逻辑
		c.Set("name","zlx") //设置上下文值,方便传递
		c.Next()
		//是否登录
		}else {
			c.Next()
		}
	}
}

func main() {
	//	中间件
	r := gin.Default()
	//gin.New() ,这个没有默认的中间件
	r.Use(midHandler,m2,autoMidlleware(true)) // 全局定义中间件

	r.GET("/index", midHandler, indexHandler,test)







	//路由组 中间件
	g1 := r.Group("/xx", autoMidlleware(true))
	{
		g1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg":"g1",
			})
		})
	}

	//路由组 中间件
	g2 := r.Group("/x2")
	g2.Use(autoMidlleware(true))
	{
		g2.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg":"g2",
			})
		})
	}



	r.Run(":9090")

}
