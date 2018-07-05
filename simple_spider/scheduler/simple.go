package scheduler

import (
	"learngo/simple_spider/engine"
)

type SimpleScheduler struct {
	Workerchan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.Workerchan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.Workerchan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		if val := reqMap[r.Url]; val {
			//log.Printf("request: %v request duplicatied\n", r.Url)
			return
		} else {
			reqMap[r.Url] = true
			s.Workerchan <- r
		}
	}()
}
