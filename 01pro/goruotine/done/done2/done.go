package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWorker(id int, w worker) {
	for v := range w.in {
		fmt.Printf("线程%d收到了%c\n", id, v)
		w.wg.Done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var ws [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		// 创建线程并且返回chan,10条线程对应10个chan
		ws[i] = createWorker(i, &wg)
	}

	for i := 0; i < 10; i++ {
		ws[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		ws[i].in <- 'A' + i
	}
	wg.Wait()

}

type person struct {
	age int
}

func (p *person) setAge(age int) {
	p.age = age
}

func main() {
	// 方法那边的self是指针还是值取决与方法
	// p := person{}
	// p.setAge(10)
	// fmt.Println(p)
	// 接口中的self如果是指针就只能接受指针类型对象,如果是值两者都可以接受
	// 接口的肚子里有个类型,还有一个指向指针指向具体数据(如果是指针类型指向的就是数据的地址&{xxx 1m0s})

	chanDemo()

}
