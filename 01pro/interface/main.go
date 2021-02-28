package main

import "fmt"

//People is xxx
type People interface {
	Speak(string) string
}

// Student is yyy
type Student struct{}

// Speak is qqq
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {

	// 不能通过编译
	// var peo People = Student{}
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
