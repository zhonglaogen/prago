package main

import (
	"fmt"
	myutil "math/pub"
)

// 最长公共子序列的问题
// 例如：  a,b,d,e    a,d,e,c 此时最长公共子序列为ade
// 时间复杂度为 2的n次方 极端情况 aaaaa  aaaaa  此时子序列为a的所有子序列，即有a没有a类似二进制
func main() {

	var s1, s2 = []rune("$abec"), []rune("$0adabe")
	fmt.Println(len(s1), len(s2))

	const count = 1000
	var f [count][count]int

	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			f[i][j] = myutil.ChooseMax(f[i-1][j], f[i][j-1])
			if s1[i] == s2[j] {
				f[i][j] = myutil.ChooseMax(f[i][j], f[i-1][j-1]+1)
			}
		}
	}

	fmt.Println(f[len(s1)-1][len(s2)-1])

}
