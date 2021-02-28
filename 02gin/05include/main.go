package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数
	//要么只有一个返回值,要么两个返回值第二个必须为error

	kua := func(name string) (string, error) {
		return name + "shuai", nil
	}
	//定义模版

	//解析模版,名字至少要和文件名相同
	t := template.New("f.tmpl")
	//传递函数
	t.Funcs(template.FuncMap{
		"kua99": kua,
	})
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse is error : %v", err)
		return
	}
	//渲染模版
	t.Execute(w, "zzz")

}

func main() {
	//模版嵌套
	http.HandleFunc("/hello", f1)
	http.HandleFunc("/tmpl", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("start is err %v", err)
		return
	}
}
func demo1(writer http.ResponseWriter, request *http.Request) {
	//	定义模版

	// 解析模版
	//先写父亲,在写孩子
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("erros is %v", err)
		return
	}
	//渲染模版
	t.Execute(writer,"zlx")
}
