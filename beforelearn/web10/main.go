package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//方法一：使用map
		/*data:= map[string]interface{}{
			"name":"Leeson",
			"message":"from Cambridge",
			"age":24,
		}
		*/
		//方法一：gin.H
		data := gin.H{
			"name":    "Leeson",
			"message": "from Cambridge",
			"age":     24,
		}

		c.JSON(http.StatusOK, data)
	})

	//方法二：结构体 灵活使用tag来做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}

	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			"fhq",
			"lab",
			999999,
		}

		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
