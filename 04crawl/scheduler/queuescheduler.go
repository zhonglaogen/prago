package scheduler

import "04crawl/engine"

type QueueScheduler struct {
	// 任务通道
	requestChan chan engine.Request
	// 工人通道,工人空闲调度器需要知道
	workerChan chan chan engine.Request
}

func (s *QueueScheduler) WorkChan() (chan engine.Request) {
	return make (chan engine.Request)
}

// 添加任务
func (s *QueueScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

//func (s *QueueScheduler) configureWorkChan(chan engine.Request) {
//	panic("implement me")
//}

// 将工人通道传递过来
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {

	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {

		// 任务队列
		var requestQ []engine.Request
		// 工人队列
		var workQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWorker <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]


			}
		}
	}()
}
