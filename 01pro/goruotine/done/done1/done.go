package main

import (
	"fmt"
)

type worker struct {
	in  chan int
	out chan bool
}

func doWorker(id int, w worker) {
	for v := range w.in {
		fmt.Printf("线程%d收到了%c\n", id, v)
		// 发生死锁的原因 第一次会阻塞在这里等着收out,第二次在传值给chan的时候,这根线程仍然被阻塞
		// w.out <- true
		// 又开了一根线程,不会阻塞此线程了
		go func() {
			w.out <- true
		}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in:  make(chan int),
		out: make(chan bool),
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var ws [10]worker
	for i := 0; i < 10; i++ {
		// 创建线程并且返回chan,10条线程对应10个chan
		ws[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		ws[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		ws[i].in <- 'A' + i
	}

	// 为了知道结果,知道所有goroutine结束,  接受为额外开的20根线程传的值
	for _, w := range ws {
		<-w.out
		<-w.out
	}

}

func main() {
	chanDemo()
}
