package main

import (
	"fmt"
	"strings"
)

func main() {
	// 返回将 s 前后端所有 cutset 包含的 utf-8 码值都去掉的字符串。
	str := "abcaacabb"
	str = strings.Trim(str, "ab")
	fmt.Println(str)


}
