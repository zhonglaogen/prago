package engine

import (
	"04crawl/fetcher"
	"log"
)

type Scheduler interface {
	Submit(request Request)
	//configureWorkChan(chan Request)
	WorkerReady(chan Request)
	WorkChan() (chan Request)
	Run()
}

type ConcurrentEngine struct {
	Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	// 所有的worker共享一个in 和 out, submit提交给了一个in
	//e.configureWorkChan(in)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Submit(r)
	}

	for {
		// 消费2
		result := <-out

		itemCount := 0
		for _, item := range result.Items {
			//log.Printf("outchan get item %d %s", itemCount, item)
			go func() {e.ItemChan<-item}()
			itemCount++
		}

		for _, request := range result.Requests {
			// 生产1, 在等待生产2执行完毕,消费1不能进行
			e.Submit(request)
		}
	}

}

// 创建工人,进行工作
func createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			// 通知调度器工人准备好了
			s.WorkerReady(in)

			// 消费1
			request := <-in

			reslut, err := doWorker(request)

			if err != nil {
				continue
			}
			// 生产2,在等待消费2执行完毕
			out <- reslut

		}
	}()
}

// 进行工作(获取body,执行解析)
func doWorker(request Request) (ParseResult, error) {
	bytes, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch Error : %s", request.Url)
		return ParseResult{}, err
	}
	return request.ParseFunc(bytes), nil
}
