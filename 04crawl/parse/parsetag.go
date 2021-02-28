package parse

import (
	"04crawl/engine"
	"regexp"
)

//const regexStr  = `<div class="([^"]+)" id="([^"]+)">`
const regexStr  = `<a href="/book/([^"]+)" target="_blank">(.*)</a>`
//解析body的内容
func ParseContent(content []byte) engine.ParseResult  {
	//<div class="new_left" id="gnleft">
	//()获得子集,也就是匹配到的元素
	compile := regexp.MustCompile(regexStr)

	submatch := compile.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
//将解析出的内容存入结果集
	for _, m := range submatch {
		result.Items = append(result.Items,m[1])
		result.Requests = append(result.Requests,engine.Request{
			Url:"https://www.dushu.com/book/" + string(m[1]),
			ParseFunc:engine.NilParse,
		})
	}
	return result

}