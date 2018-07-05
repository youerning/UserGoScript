package engine

import "learngo/simple_spider/model"

type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []model.Item
}
