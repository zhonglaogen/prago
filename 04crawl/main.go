package main

import "fmt"

var max  = map[int]int {1:7, 2:50, 3:100}

var cur  = map[int]int {1:1, 2:10, 3:66}
func main() {
	//sa := scheduler.QueueScheduler{}
	//e := engine.ConcurrentEngine{
	//	//这里为什么只能用指针,接口特性
	//	&sa,
	//	10,
	//	persist.ItemSave(),
	//}
	//e.Run(engine.Request{
	//	//Url:"https://www.dushu.com/",
	//	Url:       "https://www.dushu.com/showbook/124197/",
	//	ParseFunc: parse.ParseBookList,
	//})
	// 填充基础空间
	var left int
	left = baseArmy()
	army(50)
	fmt.Println(cur)

}
func baseArmy() int {
	baseSpace := make(map[int]int)
	maxCount := 0
	maxIndex := -1
	for e := range cur {
		if maxCount < cur[e] {
			maxCount = cur[e]
			maxIndex = e
		}
		baseSpace[e] = max[e] - cur[e]
	}


}


func army(n int) {
	if n == 0 {
		return
	}
	count := 0;
	for e := range cur {
		if cur[e] < max[e] {
			count++;
		}
	}
	if count == 0 {
		return
	}
	avg := n / count
	left := n % count
	for e := range cur {
		if cur[e] >= max[e] {
			continue
		}
		if cur[e] + avg > max[e] {
			left += max[e] - cur[e] + avg
			cur[e] = max[e]

		} else {
			cur[e] += avg
		}
	}
	if left == 0 {
		return
	} else {
		army(left)
	}

}
