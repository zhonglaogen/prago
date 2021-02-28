package main

import "fmt"

func main() {

	// fmt.Println(anagram())
	// sameNumber()
	sameNumber2()

}

//anagram  两个字符串，字母相同，但是顺序不同
func anagram() int {
	//类似计数排序，通过ascii码表来存储每一个字符的位置
	var st1 [26]int
	var st2 [26]int

	s1 := "cats"
	s2 := "act"

	for _, v := range []byte(s1) {
		st1[v-97]++
	}
	for _, v := range []byte(s2) {
		st2[v-97]++
	}

	if st1 == st2 {
		return 1
	}

	return 0

}

// 数组内的数字中有个数的相同的数量大于数组的一半，求这个数字
func sameNumber() {

	// 第一种方法，用栈
	array := [...]int{1, 3, 1, 1, 2, 1, 4}
	stack := [len(array)]int{0}
	top := -1

	for _, v := range array {
		if top == -1 {
			// 入栈
			top++
			stack[top] = v
		} else if v == stack[top] {
			top++
			stack[top] = v
		} else {
			top--
		}
	}

	fmt.Println(stack[top])

}

func sameNumber2() {

	// 第二种方法，用两个变量来替代栈
	array := [...]int{1, 3, 1, 1, 2, 1, 4}
	currentValue := 0
	count := 0

	for _, v := range array {
		if count == 0 {
			currentValue = v
			count++
		} else if currentValue == v {
			count++
		} else {
			count--
		}
	}

	fmt.Println(currentValue)

}
