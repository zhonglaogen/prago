package main

import (
	"fmt"
	"sort"
)

func generate(nodes *[]int) {
	for i := 0; i < 20; i++ {
		*nodes = append(*nodes, i)
	}
	fmt.Println(*nodes)
}
func main() {
	// 切片只能保留后面的前面的无法保留
	a := [5]int{1, 2, 3, 4, 5}
	// 定义切片
	b := []int{}
	var c []int
	d := make([]int, 10, 20)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	fmt.Println(len(d), cap(d))

	fmt.Println("==========================")

	// 只有第一个参数可以省略,cap = max - low
	b = a[1:3:4]
	fmt.Println(len(b), cap(b))
	fmt.Println("==========================")

	// copy函数的len必须足够长，不是cap
	copy(b, []int{1, 3})
	fmt.Println(b)
	fmt.Println("==========================")

	// 插入
	b = append(b[:1], append([]int{2}, b[1:]...)...)
	fmt.Println(b)
	fmt.Println("==========================")

	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	fmt.Println("==========================")

	var aa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, fmt.Sprintf("%v", i))
	}
	fmt.Println(aa)

	var a2 = [...]int{3, 7, 8, 9, 1}
	// 生序排序
	sort.Ints(a2[:])
	// 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(a2[:])))
	fmt.Println(a2)

	fmt.Println("------------------------")

	var nodes []int
	generate(&nodes)
	fmt.Println(nodes)

}
