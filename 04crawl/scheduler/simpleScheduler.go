package scheduler

import "04crawl/engine"


type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	return
}

func (s *SimpleScheduler) WorkChan() (chan engine.Request) {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 解决循环等待死锁问题
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) configureWorkChan(c chan engine.Request) {
	s.workerChan = c
}


