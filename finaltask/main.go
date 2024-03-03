package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:wxe13867187633@tcp(127.0.0.1:3306)/fhq?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

//根据表列名定义结构体
type User struct {
	Id   int
	Age  int
	Name string
}

// 查询单条数据示例
func queryRowDemo(id string) (User, error) {
	sqlStr := "select id, name, age from `user` where id=?"
	var u User
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return User{}, err
	}
	return u, nil
}

// 插入数据
func insertRowDemo(id, name, age string) {
	sqlStr := "insert into `user`(id,name,age) values (?,?,?)"
	ret, err := db.Exec(sqlStr, id, name, age) //后俩就是放变量的地方
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo(id, name, age string) {
	sqlStr := "update `user` set age=?,name=? where id = ?"
	ret, err := db.Exec(sqlStr, age, name, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(id string) bool {
	sqlStr := "delete from `user` where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return false
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return false
	}
	if n == 0 {
		fmt.Printf("delete failed, student with ID %s does not exist\n", id)
		return false
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
	return true
}


func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Printf("successfully link!")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	//加载静态文件
	r.Static("/statics", "./finaltask/statics")

	// 添加跨域配置  调整的时候多加了本来是拿来处理change路由fetch不了的问题 结果是表名忘带了
	// 添加 CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	//上传主路由的表格数据
	r.GET("/", func(c *gin.Context) {
		rows, err := db.Query("SELECT id,name,age FROM `user`")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.Id, &user.Name, &user.Age)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)

			/*
				c.HTML(http.StatusOK,"finaltask/index.html",nil)
			*/
		}

		tmpl := template.Must(template.ParseFiles("finaltask/index.html"))
		err = tmpl.Execute(c.Writer, users)
		if err != nil {
			log.Fatal(err)
		}

	})

	// GET 请求处理 /add.html 页面
	r.GET("/add.html", func(c *gin.Context) {
		r.LoadHTMLFiles("./finaltask/add.html")
		c.HTML(http.StatusOK, "add.html", gin.H{
			"title": "Add Page",
		})
	})
	// GET 请求处理 /update.html 页面
	r.GET("/update.html", func(c *gin.Context) {

		r.LoadHTMLFiles("./finaltask/update.html")
		c.HTML(http.StatusOK, "update.html", gin.H{
			"title": "Update form",
		})

	})

	//增加
	r.POST("/login", func(c *gin.Context) {
		// 解析表单数据
		err := c.Request.ParseForm()
		if err != nil {
			c.String(400, "Failed to parse form")
			return
		}

		// 获取表单数据
		ID := c.PostForm("id")
		Name := c.PostForm("name")
		Age := c.PostForm("age")

		insertRowDemo(ID, Name, Age)

		// 插入数据完成后重定向到主页
		c.Redirect(http.StatusFound, "/") // 重定向到主页或其他页面
	})

	//查询
	r.GET("/question", func(c *gin.Context) {
		// 从 URL 参数中获取用户输入的学号
		query := c.Query("query")

		// 调用查询函数并将结果传递给前端页面
		user, err := queryRowDemo(query)
		if err != nil {
			fmt.Printf("We can't find the %v student", query) //没有找到学号xxx的学生
			// 处理错误
		}

		r.LoadHTMLFiles("./finaltask/change.html")
		// 将查询结果传递到前端页面
		c.HTML(http.StatusOK, "change.html", gin.H{
			"user": user,
		})

		// 查询数据完成后重定向到主页
		/*c.Redirect(http.StatusFound, "/") // 重定向到主页或其他页面*/
	})

	//来自主路由的删除
	r.POST("/delete", func(c *gin.Context) {
		var requestData map[string]interface{}

		err := c.BindJSON(&requestData) // 将 HTTP 请求中的 JSON 数据绑定到 requestData 变量中
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
			return
		}

		ID, ok := requestData["id"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的学生ID"})
			return
		}

		fmt.Println("将要删除的学生ID为:", ID) // 打印即将删除的学生ID

		success := deleteRowDemo(ID) //进行删除

		if success {
			c.JSON(http.StatusOK, gin.H{"message": "学生删除成功"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "学生ID不存在"})
		}
	})

	//来自change小页面的删除键处理
	r.POST("/changedelete", func(c *gin.Context) {
		//解析change页面
		r.LoadHTMLFiles("./finaltask/change.html")
		fmt.Println("接收到 /changedelete 请求")

		var requestData map[string]interface{}

		err := c.BindJSON(&requestData) // 将 HTTP 请求中的 JSON 数据绑定到 requestData 变量中
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
			return
		}

		ID, ok := requestData["id"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的学生ID"})
			return
		}

		fmt.Println("将要删除的学生ID为:", ID) // 打印即将删除的学生ID

		success := deleteRowDemo(ID) //进行删除

		if success {
			c.JSON(http.StatusOK, gin.H{"message": "学生删除成功"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "学生ID不存在"})
		}

	})

	//修改按键
	r.POST("/update", func(c *gin.Context) {
		// 解析表单数据
		fmt.Println("接收到update请求")

		/*
			err := c.Request.ParseForm()
			if err != nil {
				c.String(400, "Failed to parse form")
				return
			} */

		fmt.Println("URL:", c.Request.URL)

		// 获取表单数据
		ID := c.PostForm("id") 
		Name := c.PostForm("name")
		Age := c.PostForm("age")

		fmt.Println("ID:", ID)
		fmt.Println("Name:", Name)
		fmt.Println("Age:", Age)

		updateRowDemo(ID, Name, Age)

		// 插入数据完成后重定向到主页
		c.Redirect(http.StatusFound, "/") // 重定向到主页或其他页面

	})

	r.Run(":9000")
}
