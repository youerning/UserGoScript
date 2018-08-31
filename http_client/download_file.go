// download file
package main

import (
	"net/http"
	"log"
	URL "net/url"
	"path"
	"os"
	"io"
	"fmt"
)

func main() {
	url := "<file_url>"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("download file failed\nurl: ", url)
	}

	// 必须关闭Body, 不然tcp连接不能复用
	defer resp.Body.Close()

	parsedUrl, err := URL.Parse(url)
	if err != nil{
		log.Fatal("parse url failed\nurl: ", url)
	}

	filename := path.Base(parsedUrl.Path)

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal("create file failed")
	}
	_, err = io.Copy(file, resp.Body)
	fmt.Printf("download file: %s success!", filename)
}
