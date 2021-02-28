package main

import (
	"fmt"
	myutil "math/pub"
)

func main() {

	var N, M = 4, 5
	var v = [...]int{0, 1, 2, 3, 4}
	var w = [...]int{0, 2, 4, 4, 5}
	var f [6]int

	for i := 1; i <= N; i++ {
		for j := v[i]; j <= M; j++ {
			f[j] = myutil.ChooseMax(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Println(f[M])

}
