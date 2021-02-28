package main

import (
	"fmt"
	myutil "math/pub"
)

// 时间复杂度n的三次方，开始枚举的化应该是n!
// 类似矩阵相乘，相邻的石头可以相加，求最小相加次数
func main() {

	const max = 310
	// 输入元素的个数
	n := 4
	// 和数组，存储前n个元素的和
	var sum [max]int
	// 第一维是i左侧位置，第二维是j右侧的位置,注意这个数组构建是斜着构建的从左到右从上倒下
	var f [max][max]int

	fmt.Println("请输入每一堆石头 的个数，默认为4 堆")
	for i := 1; i <= n; i++ {
		fmt.Scanln(&sum[i])
		sum[i] += sum[i-1]
	}

	for len := 2; len <= n; len++ {
		for i := 1; i+len-1 <= n; i++ {
			j := i + len - 1
			f[i][j] = 1e8
			for k := i; k < j; k++ {
				f[i][j] = myutil.ChooseMin(f[i][j], f[i][k]+f[k+1][j]+sum[j]-sum[i-1])
			}
		}

	}

	fmt.Println("结果为：", f[1][n])

}
