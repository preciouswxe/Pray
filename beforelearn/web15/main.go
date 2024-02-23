package main

import (
	
	"net/http"
	"path"
	

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
 
	r.LoadHTMLFiles("./web15/index.html")

	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/upload",func(c *gin.Context){
		//从请求中读取文件
		f,err := c.FormFile("f1")   //从请求中获取携带的参数一样的思路
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error" : err.Error(),
			})
		}else{
			//将读取到的文件保存在本地（服务端本地）
			//dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./",f.Filename) //找个路径
            c.SaveUploadedFile(f,dst)      //把客户端传上来的文件放到这个路径下面
			c.JSON(http.StatusOK,gin.H{
				"status":"OK",
			})
		}
		
	})
	r.Run(":9090")
}