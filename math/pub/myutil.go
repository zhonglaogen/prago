package myutil

//ChooseMax 寻找最大值
func ChooseMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//ChooseMin 寻找最小值
func ChooseMin(a, b int) int {
	if a < b {
		return a
	}
	return b

}
