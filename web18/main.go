package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") //获取set数据
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ......")

	//计时  统计耗时
	start := time.Now()

	//go funXX(c.Copy()) //goroutine在funcXX中只能使用c的拷贝
	
	c.Next() //责任链模式 调用后续的处理函数也就是indexHandler
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)

	fmt.Printf("cosgt:%v\n", cost)
	fmt.Println("m1 out......")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ......")
	c.Set("name", "Leeson")    //上下文中命名，让后续函数得到name 实现跨中间件
	//c.Next()  //调用后续处理函数 打印index
	//c.Abort() //阻止后续的处理函数
	//return
	fmt.Println("m2 out......")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//链接数据库
	//或者一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if是登录用户 则c.Next() 否则else c.Abort()
			c.Next()
		} else {
			c.Next()
		}

	}

}

func main() {
	r := gin.Default()

	r.Use(m1, m2, authMiddleware(true)) //注册全局中间件 作用在下面所有路由  后续就不用写m1了
	//本质会变成递归 m1调动m2，m2调取index函数，返回index'然后再打印m2结束，最后m1结束，打印m1结束

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})


	/* 
	//路由组注册中间件方法1：
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}

	//路由组注册中间件方法2：
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	} */

	r.Run(":9090")
}
