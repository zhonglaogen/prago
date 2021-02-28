package main

import "fmt"

func main() {
	// 给定一个字符串，要求把字符串前面的若干个字符移动到字符串的尾部，
	//如把字符串“abcdef”前面的2个字符'a'和'b'移动到字符串的尾部，使得原字符串变成字符串“cdefab”
	//请写一个函数完成此功能，
	//要求对长度为n的字符串操作的时间复杂度为 O(n)，空间复杂度为 O(1)。

	arrys := []rune("abcdef")
	length := len(arrys)
	// 暴力移动,
	// leftRotateString(arrys, length, 2)
	// 三步反转
	leftRotateString2(arrys, length, 2)
	fmt.Println(string(arrys))

}

func leftRotateString2(arrys []rune, n int, m int) {
	m %= n //等价于移动0位
	reverseString(arrys, 0, m-1)
	reverseString(arrys, m, n-1)
	reverseString(arrys, 0, n-1)

}

func reverseString(arrys []rune, from, to int) {
	for from < to {
		tmp := arrys[from]
		arrys[from] = arrys[to]
		from++
		arrys[to] = tmp
		to--
	}
}

// 暴力移动法
func leftRotateString(arrys []rune, length int, move int) {
	for move != 0 {
		leftShiftOne(arrys, length)
		move--
	}
}

func leftShiftOne(arrys []rune, length int) {
	var first rune = arrys[0]
	for i := 1; i < length; i++ {
		arrys[i-1] = arrys[i]
	}
	arrys[length-1] = first
}
