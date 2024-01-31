package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数kua
	//要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "他好帅啊他好帅", nil
	}

	//定义模板
	t := template.New("f.tmpl")
	//告诉模板引擎，我现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	

	//解析模板
	_, err := t.ParseFiles("./web06/f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}

	//渲染模板
	name := "腻电"

	t.Execute(w, name)
}

func demo1(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err := template.ParseFiles("./web06/t.tmpl","./web06/ul.tmpl")
	if err != nil{
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	//渲染模板
    name := "腻电2号"
	t.Execute(w,name)
}

func main() {
	http.HandleFunc("/", f1)
    http.HandleFunc("/tmplDemo",demo1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
