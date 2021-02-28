package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//text 和html包的区别,传递函数
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("start is err %v", err)
		return
	}
}
func xss(writer http.ResponseWriter, request *http.Request) {
	//	定义模版
	//解析模版
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(name string) template.HTML {
			return template.HTML(name)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("parse is error")
		return
	}

	//渲染模版
	str := "<script>alert(123)</script>"
	str2 := "<a href='https://www.baidu.com'> 百度 </a>"
	t.Execute(writer, map[string]string{
		"str1": str,
		"str2" : str2,
	})
}
func index(writer http.ResponseWriter, request *http.Request) {
	//定义模版
	//	解析模版
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse is error")
		return
	}

	//渲染模版
	t.Execute(writer, "zlx")

}
