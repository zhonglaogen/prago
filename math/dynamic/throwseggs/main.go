package main

import (
	"fmt"
	myutil "math/pub"
)

const (
	//N 是最高高度
	N = 110
	//M 是最多鸡蛋个数
	M = 11
)

func main() {
	throw2()
}

func throw1() {
	var n, m int
	// 属性代表次数，i为高度j为鸡蛋个数
	var f [N][M]int
	fmt.Println("请输入高度(100) 和 鸡蛋的个数(2)---（14）：")
	fmt.Scanf("%d %d", &n, &m)

	for i := 1; i <= n; i++ {
		f[i][1] = i
	}
	for i := 1; i <= m; i++ {
		f[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			f[i][j] = f[i][j-1]
			// k为在哪里选择落下去
			for k := 1; k <= i; k++ {
				f[i][j] = myutil.ChooseMin(f[i][j], myutil.ChooseMax(f[k-1][j-1], f[i-k][j])+1)
			}
		}

	}

	fmt.Println(f[n][m])
}

func throw2() {
	var n, m int
	// 属性代表高度，i为次数j为鸡蛋个数
	var f [N][M]int
	fmt.Println("请输入高度(100) 和 鸡蛋的个数(2)---（14）：")
	fmt.Scanf("%d %d", &n, &m)

	//此种方式永远都会有解，极端情况n层楼一个蛋 测n此也会出结果

over:
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = f[i-1][j-1] + f[i-1][j] + 1
			if f[i][j] >= n {
				fmt.Println(i)
				break over
			}
		}
	}

}
