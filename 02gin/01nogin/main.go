package _1nogin

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello",sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		fmt.Printf("error :%v\n",err)
		return
	}
}
func sayhello(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer,"<h1>hello</h1>\n")
	bytes, e := ioutil.ReadFile("./hello.txt")
	if e != nil {
		fmt.Printf("error %v\n",e)
	}
	fmt.Fprint(writer,string(bytes))

}


