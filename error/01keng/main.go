package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// 结果显示为true，如果大括号写在下面，就会默认在switch的后面加上分号，为switch的加上默认值为true
	// testSwitch()

	// 不会执行defer的内容
	// testOsDefer()

	// 会执行，但是报错了，会中断goroutine
	// testGoDefer()

	//位运算变量和常量不同 结果是 （推断出1是byte）0溢出byte， 24，（推断1为不确定类型）常量为不确定类型
	// testVar()

	// 报错，多级表达式没有前后顺序
	// testSlice()

	// range遍历是赋值给一个值的方式....
	// unsafe.Sizeof(int(0))

	// student{1}.fmtPointer()

}

type student struct {
	age int
}

func (Lili *student) fmtPointer() {
	fmt.Println("poniter")
}

func (Lili student) fmtReference() {
	fmt.Println("reference")
}

func testSlice() {
	var m []int = nil

	// m = []int{1, 2}
	// m[0] = 8

	// 编译成功，运行失败，提示index out of range [0] with length 0，所以同时运行
	m, m[0] = []int{1, 2}, 8

	fmt.Println(m)
}

func testVar() {

	var n uint = 10
	const N uint = 10

	var x byte = (1 << n) % 100
	var y byte = (1 << N) % 100
	fmt.Printf("%d, %d", x, y)
}

func testGoDefer() {
	defer fmt.Println("hello")
	runtime.Goexit()
}

func testOsDefer() {
	defer fmt.Println("hello")
	os.Exit(1)
}

func testSwitch() {
	switch p(); {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	}
}

func p() bool {
	return false
}
