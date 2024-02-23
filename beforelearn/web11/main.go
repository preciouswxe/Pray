package main

//querystring
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// GET请求 URL ？后面的是querystring参数
	// key=value格式，多个key-value用 & 链接
	// eq : /web?query=xxxxxx&age=xxx
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的 query string 参数
		name := c.Query("query")
		age := c.Query("age")
		//通过Query获取请求中携带的querystring参数
		//name := c.DefaultQuery("query","somebody")
		/*name,ok := c.GetQuery("query")
				if !ok {
		            //取不到
					name ="somebody"
				}*/

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}
