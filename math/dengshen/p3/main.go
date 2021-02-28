package main

import (
	"fmt"
)

// 黑袍
func main() {
	isPowerOfTwo1(8192)
	isPowerOfTwo1(0)
	isPowerOfTwo2(8192)
	isPowerOfTwo2(0)
	// isPowerOfTwo3(8192)
	isPowerOfTwo3(0)
	isPowerOfTwo4(8192)
	isPowerOfTwo4(1)
}

// 判断给出的整数是否是2的n次方
func isPowerOfTwo1(n int) {
	if n == 0 {
		fmt.Println("不是2的幂")
		return
	}
	for n%2 == 0 {
		n = n / 2
	}
	if n == 1 {
		fmt.Println("是2的幂")
		return
	}
	fmt.Println("不是是2的幂")
}

// 击败63
func isPowerOfTwo2(n int) {
	if n <= 0 {
		fmt.Println("不是是2的幂")
		return
	}
	if ((n - 1) & n) == 0 {
		fmt.Println("是2的幂")
		return
	}
	fmt.Println("不是是2的幂")
}

// 如果是2的幂（仅有一位是1），与他的相反数（按位取反+1）只有1的位和后面的位是相同的
// 击败63%
func isPowerOfTwo3(n int) {

	if n <= 0 {
		fmt.Println("false")
		return
	}

	if n&(-n) == n {
		fmt.Println("ok")
		return
	}
	fmt.Println("false")
}

// 最大的2的幂除以2的幂都为 0，，，，击败100%
func isPowerOfTwo4(n int) {

	if n <= 0 {
		fmt.Println("false")
		return
	}

	if (1<<30)%n == 0 {
		fmt.Println("ok")
		return
	}
	fmt.Println("false")
}
