package main

import (
	"fmt"
	myutil "math/pub"
)

//COUNT 常量值
const COUNT = 1000

func main() {
	goodBag()
}

//  优化后的01背包
func goodBag() {

	var N, M = 4, 5
	var v = [...]int{0, 1, 2, 3, 4}
	var w = [...]int{0, 2, 4, 4, 5}
	var f [6]int

	// var M, N int
	// var v [COUNT]int
	// var w [COUNT]int
	// var f [COUNT][COUNT]int

	// // fmt.Println("请输入物品数量和背包的容量")
	// fmt.Scanf("%d %d", &N, &M)
	// // fmt.Println("输入物品的的体积和价值")
	// for i := 1; i <= N; i++ {
	// 	fmt.Scanf("%d %d", &v[i], &w[i])
	// }
	for i := 1; i <= N; i++ {
		for j := M; j >= v[i]; j-- {
			f[j] = myutil.ChooseMax(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Println(f[M])
}

// 未经优化的01背包
func simBag() {
	// var N, M = 4, 5
	// var v = [...]int{0, 1, 2, 3, 4}
	// var w = [...]int{0, 2, 4, 4, 5}
	// var f [5][6]int

	var M, N int
	var v [COUNT]int
	var w [COUNT]int
	var f [COUNT][COUNT]int

	// fmt.Println("请输入物品数量和背包的容量")
	fmt.Scanf("%d %d", &N, &M)
	// fmt.Println("输入物品的的体积和价值")
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d %d", &v[i], &w[i])
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = myutil.ChooseMax(f[i-1][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}

	fmt.Println(f[N][M])
}
