/*
 * @Author: 13561 1356137745@qq.com
 * @Date: 2024-02-19 16:29:03
 * @LastEditors: 13561 1356137745@qq.com
 * @LastEditTime: 2024-02-19 17:35:46
 * @FilePath: \代码！\web13\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取请求的path（URI）参数 ，返回的都是字符串类型
//注意url的匹配不要冲突

func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(c *gin.Context) {
		//获取路径参数
		name := c.Param("name")
		age := c.Param("age") //string类型

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.Run(":9090")
}
