package engine

import (
	"log"
	"fmt"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		result, err := Worker(req)
		if err != nil {
			fmt.Println("worker error: ", err)
		}
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Item: %v\n", item)
		}
	}

}
