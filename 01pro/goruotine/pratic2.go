package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	wgs sync.WaitGroup
)

func generater(jobChan chan<- int64) {
	// defer wgs.Done()
	for i := 0; i < 24; i++ {
		r := rand.Int63()
		jobChan <- r
	}
	close(jobChan)

}

func randSum(num int64) int64 {
	var sum int64
	for num > 0 {
		x := num % 10
		sum += x
		num /= 10
	}
	return sum
}

func worker(jobChan <-chan int64, res chan<- int64) {
	// defer wgs.Done()
	for job := range jobChan {
		sum := randSum(job)
		res <- sum
		fmt.Println("收到的数字为:", sum)
	}

}

func main() {

	jobChan := make(chan int64)
	// 加上缓冲原因就是,每个worker都给这个结果chan发结果,导致线程没结束,没有done掉,所有线程都等待发,主线程wait也没法收,死锁
	// 不加wait,但是加上close会导致主线程提前关闭chan,其他线程发送结果报错send异常
	// 不加wait和close 主线程一直再收,没有关闭通道,没有其他线程发
	// 总结:其他线程可以是死循环的收,主线程不可以.因为其他线程一直收,主线程结束其他线程就结束;但是主线程一直收,没办法停下来
	// close的chan再次close会报错
	// close的chan发送 会报错
	// close的chan可以继续接受数据
	resultChan := make(chan int64, 100)
	// wgs.Add(1)
	go generater(jobChan)

	for i := 0; i < 24; i++ {
		// wgs.Add(1)
		go worker(jobChan, resultChan)
	}

	// wgs.Wait()
	// close(resultChan)
	for v := range resultChan {

		fmt.Println(v)
	}

}