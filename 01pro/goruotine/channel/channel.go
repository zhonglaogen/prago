package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for v := range c {
		fmt.Printf("线程%d收到了%c\n", id, v)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channles [10]chan<- int
	for i := 0; i < 10; i++ {
		// 创建线程并且返回chan,10条线程对应10个chan
		channles[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channles[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channles[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
