package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./beforelearn/web08/index.tmpl")

	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
	}
	//渲染模板
	name := "寒假第一战"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("execute template failed,err:%v\n", err)
		return
	}

}

func xss(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板
	//解析模板之前定义一个自定义的函数safe再解析
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./beforelearn/web08/xss.tmpl")

	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	//渲染模板
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://hdufhq.cn:8888/'>hdu孵化器地址</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
