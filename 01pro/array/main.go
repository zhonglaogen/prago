package main

import "fmt"

func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	// var xxx = [...]int{1, 2}
	var w1 = []int{1, 3}

	// 插入
	w1 = append(w1[:1], append([]int{2}, w1[1:]...)...)
	fmt.Println(w1)
	fmt.Println(testArray) //[0 0 0]
	fmt.Println(numArray)  //[1 2 0]
	fmt.Println(cityArray) //[北京 上海 深圳]

	// var numArray2 = [...]int{1, 2}

	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}

	var xxx = [...]int{1, 2}

	var yyy = [...]int{1, 2}

	println(xxx == yyy)

	// xb := -1
	// 相当与c的~去反操作
	fmt.Println(^-1)
	fmt.Println(^0)
	fmt.Println(^1)
	fmt.Println(^3)

}
