package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./web12/login.html","./web12/index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 处理这个访问login请求，上面只是进入到页面的get请求

	r.POST("/login", func(c *gin.Context) {
		//获取form表达的提交的数据
		/*方法一
		username := c.PostForm("username")
		password := c.PostForm("password")*/

        /*方法二  
	    username :=c.DefaultPostForm("username","somebody")
		password :=c.DefaultPostForm("password","***") */

        //方法三

        username ,ok:= c.GetPostForm("username")
		if !ok{
			username = "preciouswxe"
		}
		password ,ok:=c.GetPostForm("password") //ok被重新创建
		if !ok{
			password = "i don't know"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
