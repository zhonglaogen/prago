package parse

import (
	"04crawl/engine"
	"regexp"
	"strings"
)

const BookListRe =  `<a href="/showbook/([^"]+)">([^<]+)</a>`

// 获取body解析的内容,并返回结果(items, 新的request)
func ParseBookList(contetns []byte) engine.ParseResult {

	re := regexp.MustCompile(BookListRe)

	submatch := re.FindAllSubmatch(contetns, -1)

	result := engine.ParseResult{}

	for _, m := range submatch {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: func(joinUrl string) string {
				if strings.Contains(joinUrl, "http") {
					return joinUrl
				}
				return "https://www.dushu.com/showbook/" + joinUrl
			} (string(m[1])),
			ParseFunc: engine.NilParse,
		})
	}

	return result

}
