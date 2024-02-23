package main

import (
	
    "net/http"
	"github.com/gin-gonic/gin"
)

type UserInfo struct{
	Username string  `form:"username" json:"username"`
	Password string  `form:"password" json:"password"`
}

func main(){
	r:= gin.Default()

	r.LoadHTMLFiles("./web14/index.html")

	r.GET("/user",func(c *gin.Context){
		/* 
		username := c.Query("username")
		password := c.Query("password")
		u := UserInfo{
			username: username,
			password: password,
		} */
        var u UserInfo //声明一个UserInfo类型的变量u
	    err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是拿出c这个http请求文内的数据
	                        //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
	    if err != nil { 
     	    c.JSON(http.StatusBadRequest,gin.H{
			    "error":err.Error(),
		    })
	    }else{
		    c.JSON(http.StatusOK,gin.H{
			    "status":"ok",
		    })
	    }
	})

	r.GET("/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/form",func(c *gin.Context){
		var u UserInfo //声明一个UserInfo类型的变量u
	    err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是对照着键名拿出c这个http请求文内的数据
	                        //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
	    if err != nil { 
     	    c.JSON(http.StatusBadRequest,gin.H{
			    "error":err.Error(),
		    })
	    }else{
		    c.JSON(http.StatusOK,gin.H{
			    "status":"ok",
		    })
		}
	})

	r.POST("/json",func(c *gin.Context){
		var u UserInfo //声明一个UserInfo类型的变量u
	    err := c.ShouldBind(&u) //顺着指针来修改u的值 本质是对照着键名拿出c这个http请求文内的数据
	                        //一般是用在不知道上传什么数据 要求u的键是大写字母 写上tag
	    if err != nil { 
     	    c.JSON(http.StatusBadRequest,gin.H{
			    "error":err.Error(),
		    })
	    }else{
		    c.JSON(http.StatusOK,gin.H{
			    "status":"ok",
		    })
		}
	})


		
	r.Run(":9090")
}