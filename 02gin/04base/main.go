package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	gender string
	Age int
}


func sayhello(writer http.ResponseWriter, request *http.Request) {
	//	定义模版

	//解析模版
	temp, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse is error: %v", err)
		return
	}
	//渲染模版
	u1 := User{
		Name:"钟老根",
		gender: "man",
		Age: 20,
	}
	m1 := map[string]interface{}{
		"Name": "钟老根",
		"Gender": "man",
		"Age": 10,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"游泳",
	}
	temp.Execute(writer, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})
}


func main() {
	//传送数据给模版
	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("server start is error: %v", err)
		return
	}
}
