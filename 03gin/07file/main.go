package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	//上传文件
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	r.POST("/onfile", func(c *gin.Context) {

		f1, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":err.Error(),
			})
		}else {
			//dst := fmt.Sprintf("./%s",f1.Filename)
			dst := path.Join("./",f1.Filename)
			c.SaveUploadedFile(f1,dst)
			c.JSON(http.StatusOK,gin.H{
				"msg": "ok",
			})
		}
	})
	r.Run(":9090")
}
