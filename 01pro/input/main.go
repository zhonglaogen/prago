package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入 一个字符")
	var a string
	fmt.Scanln(&a)
	fmt.Println(a)
	keyWord()

}

func keyWord() int {
	fmt.Println("input:")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	return 1

}
