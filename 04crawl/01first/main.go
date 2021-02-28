package main

import (
	"04crawl/fetcher"
	"04crawl/parse"
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func main() {
	//解决编码问题
	//go get -u golang.org/x/text
	//go get -u golang.org/x/net/html
	resp, err := http.Get("http://www.chinanews.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error status code %d\n", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := fetcher.DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	bytes, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s", bytes)
	parse.ParseContent(bytes)

}

