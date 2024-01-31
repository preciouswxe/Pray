package main

import(
	"fmt"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter,r *http.Request){
    //定义模板
	//解析模板
	t,err := template.ParseFiles("./web07/index.tmpl")
    if err != nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	msg := "泥电index"
	//渲染模板
	t.Execute(w,msg)
}

func home(w http.ResponseWriter,r *http.Request){
    //定义模板
	//解析模板
	t,err := template.ParseFiles("./web07/home.tmpl")
    if err != nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	msg := "泥电home"
	//渲染模板
	t.Execute(w,msg)
}

func index2(w http.ResponseWriter,r *http.Request){
    //定义模板（模板继承的方式）
	//解析模板
	t,err := template.ParseFiles("./web07/templates/base.tmpl","./web07/templates/index2.tmpl")
	if err !=  nil{
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	//渲染模板
	name := "侯明昊"
	t.ExecuteTemplate(w,"index2.tmpl",name)
}

func home2(w http.ResponseWriter,r *http.Request){
    //定义模板（模板继承的方式）
	//解析模板
	t,err := template.ParseFiles("./web07/templates/base.tmpl","./web07/templates/home2.tmpl")
	if err !=  nil{
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	//渲染模板
	name := "刘宇"
	t.ExecuteTemplate(w,"home2.tmpl",name)
}

func main() {
	http.HandleFunc("/index", index)
    http.HandleFunc("/home",home)
	http.HandleFunc("/index2", index2)
    http.HandleFunc("/home2",home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
