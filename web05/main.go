package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./web05/hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:%v", err)
	}
	//渲染模板

	u1 := User{
		Name: "小王子",
		Gender: "男",
		Age: 18,
	}

	m1 := map[string]interface{}{
		"Name": "泥电",
		"Gender": "未知~~~",
		"Age": 18,
	}

	hobbyList := []string{
		"唱",
		"跳",
		"rap",
		"打篮球",
	}

	t.Execute(w, map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP sever start failed,err:%v", err)
		return
	}
}
