package engine

import (
	"log"
	"learngo/simple_spider/fetcher"
)

func Worker(req Request) (ParserResult, error) {
	log.Printf("Fetching url: %s\n", req.Url)
	resp, err := fetcher.Fetcher(req.Url)
	if err != nil {
		log.Printf("fetcher error: %v when fetch url: %s\n", err, req.Url)
		return ParserResult{}, err
	}

	return req.ParserFunc(resp), nil
}

