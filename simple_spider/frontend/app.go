package main

import (
	"net/http"
	"learngo/simple_spider/frontend/controller"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	templateFilename := "simple_spider/frontend/view/template_lower.html"
	http.Handle("/search", controller.CreateFrontendHandler(templateFilename))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
