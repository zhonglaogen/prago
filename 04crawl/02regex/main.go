package main

import (
	"04crawl/parse"
	"fmt"
	"regexp"
)

func main() {
	str := `<div class="book-title">
    <h1>女生寝室</h1>
    </div>`
	//re := regexp.MustCompile(`[0-9a-zA-Z]*@163.com`)
	re := regexp.MustCompile(parse.BookListRe)
	s := re.FindAllSubmatch([]byte(str), -1)
	//s := re.FindString(str)
	fmt.Printf("%s", s[0][1])
}
