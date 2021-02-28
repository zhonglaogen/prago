package main

import "fmt"

func main() {
	a := make(chan int)
	go func() {
		// 不关闭,主线程就会一直收,死锁
		// defer close(a)
		a <- 1
	}()

	for {
		v, ok := <-a
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
