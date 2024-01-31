package main

import (
	"fmt"
	"net/http"
	"html/template"
)


func sayHello(w http.ResponseWriter, r *http.Request){
	//先解析模板
    t , err := template.ParseFiles("./web04/hello.tmpl")   //请勿刻舟求剑
	if err != nil{
		fmt.Println("Parse template failed,err:%v",err)
		return 
	}
	//再渲染模板
	name := "你爹"
	err = t.Execute(w,name)
	if err != nil{
		fmt.Println("render template failed,err:%v",err)
		return
	}
}

func main(){
    http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil{
		fmt.Println("HTTP sever start failed,err:%v",err)
		return
	}
}