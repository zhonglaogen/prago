package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	//go get -u github.com/gin-gonic/gin
	// go mod tidy go mod download

	engine := gin.Default()

	//加载静态文件

	engine.Static("/xxx","./statics")

	//gin 框架自定义函数
	engine.SetFuncMap(template.FuncMap{
		"safe": func(name string)template.HTML {
			return template.HTML(name)
		},
	})

	//engine.LoadHTMLFiles("templates/posts/index.tmpl","templates/user/index.tmpl") //模版解析
	//gin框架批量解析文件
	engine.LoadHTMLGlob("templates/**/*")


	engine.GET("/users/index", func(c *gin.Context) {
		//name 为解析后文件的名称,或define的名字
		c.HTML(http.StatusOK,"users1/index.tmpl",gin.H{ //模版渲染,默认为文件的名字,可以用define指定
			"msg": "posts/index.tmpl",
		})
	})

	engine.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"posts1/index.tmpl",gin.H{ //模版渲染
			"msg": "<a href='https://www.baidu.com'>百度</a>",
		})
	})
	engine.Run(":9000")
}
