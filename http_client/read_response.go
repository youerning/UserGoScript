package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Status)
	fmt.Println("Headers")
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
}
