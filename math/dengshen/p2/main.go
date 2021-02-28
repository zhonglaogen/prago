package main

import (
	"fmt"
)

func main() {
	findMissNum()
	findMissNum2()
	findSingleNum()
	findSingleNum2()
}

// 查找丢失的数字，第一种方法用一个额外的数组
func findMissNum() {
	// 无序的，
	array := [...]int{0, 1, 2, 3, 5, 6}

	arrayCopy := [len(array) + 1]bool{false}
	for _, v := range array {
		arrayCopy[v] = true
	}

	for i, v := range arrayCopy {
		if v == false {
			fmt.Printf("丢失的数字为：%d\n", i)
		}
	}
}

// 查找丢失的数字，第二种，求和公式，然后减掉所有的值就是丢失的值

func findMissNum2() {
	// 无序的，
	array := [...]int{0, 1, 2, 3, 5, 6}
	len := len(array)

	sum := (len + 1) * len / 2

	for _, v := range array {
		sum -= v
	}

	fmt.Printf("丢失的数字是 %d\n", sum)
}

// 所有的数字都是两个，只有一个数字是一个，找出这个数字
func findSingleNum() {

	// 暴力方式
	array := []int{2, 3, 4, 5, 2, 3, 5}
	for _, v := range array {
		if getCount(&array, v) == 1 {
			fmt.Println("没有成对出现的数字为", v)
		}
	}
}

func getCount(array *[]int, value int) int {
	count := 0
	for _, v := range *array {
		if v == value {
			count++
		}
	}
	return count
}

// a^b^a =a 的性质求出未成对出现的 a^a = 0    0^任何数字都得到它本身
func findSingleNum2() {
	array := []int{2, 3, 4, 0, 0, 5, 2, 3, 5}
	singleNum := 0
	for _, v := range array {
		singleNum ^= v
	}
	fmt.Println("没有成对出现的数字为", singleNum)
}
