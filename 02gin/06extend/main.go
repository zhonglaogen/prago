package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//模版的使用,继承
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("start is err %v", err)
		return
	}
}


func index2(writer http.ResponseWriter, request *http.Request) {
	//定义模版(模版继承的方式)
	//解析模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("parse error")
		return
	}
	//渲染模版,加载两个模版的方式
	t.ExecuteTemplate(writer,"index2.tmpl","zlx")
}

func home2(writer http.ResponseWriter, request *http.Request) {
	//定义模版(模版继承的方式)
	//解析模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("parse error")
		return
	}
	//渲染模版,加载两个模版的方式
	t.ExecuteTemplate(writer,"home2.tmpl","zlx")
}




func home(writer http.ResponseWriter, request *http.Request) {
	//	定义模版
	//解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("parse error")
	}
	//渲染模版
	t.Execute(writer,"这是home")
}

func index(writer http.ResponseWriter, request *http.Request) {
	//	定义模版
	//解析模版
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse error")
	}
	//渲染模版
	t.Execute(writer,"这是index")
}
