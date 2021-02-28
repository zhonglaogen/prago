package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, in chan int) {
	for v := range in {
		fmt.Println("收到", v)
		time.Sleep(time.Second)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c

}

func main() {
	var g1, g2 = generator(), generator()
	var worker = createWorker(0)

	// n := 0
	// hashValue := false
	// 直接用append不用make初始化,append数组应该用...
	var values []int

	// 让程序运行10秒
	tm := time.After(10 * time.Second)
	// 定时任务
	tick := time.Tick(time.Second)

	for {
		// nil的chan不会被扫描到
		// 如果不加一个东西存储收到的n,很可能会导致n的速度过快,来不及收就下一次赋值了
		// 生产速度大于消费速度就要考虑队列存起来消费了
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-g1:
			// fmt.Println("g1", n)
			// 收到值后将值发送给worker,如果此时worker阻塞,整个select就阻塞了
			// worker <- n
			// hashValue = true
			values = append(values, n)
		case n := <-g2:
			// fmt.Println("g2", n)
			// worker <- n
			// hashValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			// hashValue = false
			values = values[1:]
			// 两次select间隔超过800毫秒就执行
		case <-tick:
			fmt.Println("长度", len(values), cap(values))
		case <-time.After(time.Millisecond * 800):
			fmt.Println("timeout")
		case <-tm:
			fmt.Println("bty")
			return
		}
	}

}
