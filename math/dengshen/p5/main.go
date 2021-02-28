package main

import "fmt"

func main() {
	aa := 017

	fmt.Printf("%#X\n", aa)
	countNumeberOne(-1)
	countNumeberOne2(-1)
	countNumeberOne2(13)

	a := reverbinary(23)
	b := reverbinary(a)
	fmt.Println(b)
}

// 给定一个二进制数字，寻找1的个数
func countNumeberOne(num int32) {
	count := 0
	for i := 0; i < 32; i++ {
		if (num>>i)&1 == 1 {
			count++
		}
	}
	fmt.Println("1的个数为:", count)
}

// 通过归并的思想，先计算两位在计算四位在计算八位
func countNumeberOne2(a int) {
	m1 := 0x55555555
	m2 := 0x33333333
	m4 := 0x0f0f0f0f
	m8 := 0x00ff00ff
	m16 := 0x0000ffff

	b := (a & m1) + ((a >> 1) & m1)
	c := (b & m2) + ((b >> 2) & m2)
	d := (c & m4) + ((c >> 4) & m4)
	e := (d & m8) + ((d >> 8) & m8)
	f := (e & m16) + ((e >> 16) & m16)

	fmt.Printf("1的个数为%d\n", f)
}

// 将二进制的八位数逆序输出，先逆序前四位和后四位相加
func reverbinary(num int) int {
	m4 := 0x0f
	m2 := 0x33
	m1 := 0x55

	num = ((num & m4) << 4) + ((num >> 4) & m4)
	num = ((num & m2) << 2) + ((num >> 2) & m2)
	num = ((num & m1) << 1) + ((num >> 1) & m1)

	return num

}
