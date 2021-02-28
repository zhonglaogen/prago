package engine

import (
	"04crawl/fetcher"
	"fmt"
	"log"
)

type simpleEngine struct {

}

func (s simpleEngine)Run(seeds ...Request) {
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		r := requests[0]
		log.Printf("url is %s",r.Url)
		requests = requests[1:]
		//获取该url的body内容
		bytes, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch err is %v", err)
		}
		//解析内容,并返回结果(获取到的url和解析出的内容),此处默认返回一个空的结果集
		parseResult := r.ParseFunc(bytes)
		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items{
			fmt.Printf("got item is %s\n",item )
		}

	}
}
