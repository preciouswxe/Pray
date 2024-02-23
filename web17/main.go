/*
 * @Author: 13561 1356137745@qq.com
 * @Date: 2024-02-22 15:14:02
 * @LastEditors: 13561 1356137745@qq.com
 * @LastEditTime: 2024-02-22 22:01:48
 * @FilePath: \代码！\web17\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//lesson17 路由与路由组

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	//访问/index的get请求走这条逻辑 会去执行下面这个匿名函数
	//用于访问得到服务器到客户端的数据 比如来获取之前存储的个人信息 个性签名 昵称之类的
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	//创建操作 比如注册会员等由客户端输入的数据 再从客户端传给服务器存储
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	//更新个人信息 部分信息 把生日修改为xx
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//删除 比如注销账号
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	//大集合 大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			//............省略
		}
		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})
	})

	//noRoute怎么处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "not found~",
		})
	})

	//视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	//路由组 也就是路由的组
	//把公用的前提提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index",func(c *gin.Context){
		    c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	    })
	    videoGroup.GET("/xx",func(c *gin.Context){
		    c.JSON(http.StatusOK,gin.H{"msg":"/video/xx"})
	    })
	    videoGroup.GET("/aaa",func(c *gin.Context){
		    c.JSON(http.StatusOK,gin.H{"msg":"/video/aaa"})
	    })
	}
	//商城的首页和详情页

	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})

	r.Run(":9090")
}
