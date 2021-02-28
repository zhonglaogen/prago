package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//使用模版
	http.HandleFunc("/hello",sayhello)
	err := http.ListenAndServe(":9000", nil)
	if err !=nil {
		fmt.Printf("error is %v",err)
		return

	}
}
func sayhello(writer http.ResponseWriter, request *http.Request) {

	//定义模版

	//解析模版
	temp, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("error is %v",err)
	}
	//渲染模版
	err = temp.Execute(writer, "钟老根")
	if err != nil {
		fmt.Println("error is %v", err)
		return
	}
	//要执行go build ,go run 不一定会编译在哪里,找不到相对路径

}
