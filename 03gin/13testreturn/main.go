package main

import (
	"fmt"
	_ "03gin/13testreturn/pack"

	)

type person struct {
	id   int
	name string
}

func main() {
	p := getNum()
	//i := getInt()
	//ss := getSlice()
	var pp *person
	pp.id=1
	fmt.Println(p,pp)
}

func getNum() (p *person) {
	//指针没有指向的内存
	p = new(person)
	p.name = "xxx"
	p.id = 1

	return
}

func getInt() (i *int) {
	//指针没有指向的内存
	*i = 10
	return
}

func getSlice() (ss *[]int) {
	i := 1
	*ss = append(*ss,i)  //这只是一个指针,没有初始化指针指向的位置
	//如果参数是 ss []*int,可以添加元素,表示的是切片,切片不初始化也可以使用

	return
}
