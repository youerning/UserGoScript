package engine

import "learngo/simple_spider/model"

type Corcurrent struct {
	Scheduler Scheduler
	Worker    int
	ItemChan chan model.Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(r Request)
	WorkerChan() chan Request

	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *Corcurrent) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()
	//e.Scheduler.ConfigureMasterWorkerChan(in)

	// create worker and ready for work
	for i := 0; i < e.Worker; i++ {
		in := e.Scheduler.WorkerChan()
		CreateWorker(in, out, e.Scheduler)
	}

	// submit request to request channel
	for _, req := range seeds {
		e.Scheduler.Submit(req)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func CreateWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}
