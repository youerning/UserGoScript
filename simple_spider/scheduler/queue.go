package scheduler

import (
	"learngo/simple_spider/engine"
)

var (
	reqMap = make(map[string]bool)
	)

type QueueScheduler struct {
	RequestChan chan engine.Request
	Workerchan  chan chan engine.Request
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(r engine.Request) {
	if val := reqMap[r.Url]; val {
		//log.Printf("request: %v request duplicatied\n", r.Url)
		return
	} else {
		reqMap[r.Url] = true
		q.RequestChan <- r
	}
}

func (q *QueueScheduler) Run() {

	q.RequestChan = make(chan engine.Request)
	q.Workerchan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWoker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWoker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-q.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-q.Workerchan:
				workerQ = append(workerQ, w)
			//send active request to active worker
			case activeWoker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}

	}()

}
func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.Workerchan <- w
}
