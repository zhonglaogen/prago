package main

import "fmt"

func main() {
	m1 := make(map[string]string, 10)
	m2 := map[string]string{
		"x1": "1",
		"x2": "2",
	}

	value, ok := m2["key"]
	if ok {
		fmt.Println(value)
	}

	fmt.Println(m1)
	fmt.Println(m2)
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	for k := range m2 {
		fmt.Println(k)
	}

	delete(m2, "x1")

	fmt.Println("===============")

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	fmt.Println(len(s), cap(s))
	s = append(s, 3)
	fmt.Println(len(s), cap(s))
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	// 结果为 1 3 3
	fmt.Printf("%+v\n", m["q1mi"])
}
