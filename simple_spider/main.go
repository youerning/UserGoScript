package main

import (
	"learngo/simple_spider/engine"
	"learngo/simple_spider/zhenai/parser"
	"learngo/simple_spider/scheduler"
	"learngo/simple_spider/persist"
)

//func main() {
//	url := "http://www.zhenai.com/zhenghun"
//
//	seed := engine.Request{
//		Url:url,
//		ParserFunc:parser.ParserCitylList,
//	}
//   e := engine.SimpleEngine{}
//	 e.Run(seed)
//}


func main() {
	url := "http://www.zhenai.com/zhenghun"
	//url := "http://www.zhenai.com/zhenghun/shanghai"

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	seed := engine.Request{
		Url:url,
		ParserFunc:parser.ParserCityList,
	}
	e := engine.Corcurrent{
		Scheduler: &scheduler.QueueScheduler{},
		Worker:30,
		ItemChan: itemChan,
	}

	//e := engine.Corcurrent{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	Worker:30,
	//}

	e.Run(seed)
}

