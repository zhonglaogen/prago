package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userInfo struct {
	//一定要大写,通过反射获取值
	Name     string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	//动态参数绑定
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {

		var u userInfo
		//c.BindJSON(&u)
		err := c.ShouldBind(&u) //??
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":8080")

}
