// make custom request for request
package main

import (
	"net/http"
	"log"
	"time"
	"fmt"
	"io/ioutil"
)

func main() {
	url := "https://www.baidu.com/"
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36"
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("new reuqest failed")
		log.Fatal(err)
	}

	request.Header.Add("User-Agent", ua)
	cookie := http.Cookie{Name:"custome_cookie",Value:"test", Expires:time.Now().Add(120 * time.Second)}

	request.AddCookie(&cookie)
	resp, err := http.DefaultClient.Do(request)

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("dump response failed")
	}
	fmt.Println(string(html))
}
