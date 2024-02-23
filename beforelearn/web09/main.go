package main

import (
	"net/http"
	
    "html/template"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//创建一个默认路由
	r := gin.Default()
    
	//解析模板

   	//加载静态文件
	r.Static("statics","/statics")

    //先给gin框架中模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe" : func(str string) template.HTML{
		    return template.HTML(str)
		},
	})


	//旧方法解析r.LoadHTMLFiles("./web09/templates/index.tmpl")

	r.LoadHTMLGlob("templates/**/*") //意思是直接加载这个web09下面的所有文件 不用具体指定路径了

	//渲染模板  在这里写上浏览器上url后小截
	r.GET("/posts/index", func(c *gin.Context) {

		//HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "<a href='http://hdufhq.cn:8888/'>孵化器地址</a>",
		})

	})
	r.GET("/users/index", func(c *gin.Context) {

		//HTTP请求要有状态码的响应被返回 三次握手 渲染上数据
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://www.hdu.edu.cn/main.htm'>杭电地址</a>",
		})

	})

    //返回从网上下载的模板,没有需要上传的数据，所以是nil
	r.GET("/home",func(c *gin.Context){
		c.HTML(http.StatusOK,"posts/home.html",nil)
	})

	//启动server 这里浏览器就是客户端 本机是服务器
	r.Run(":9090")
}
