// disable certificate check
package main

import (
	"net/http"
	"crypto/tls"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	url := "<https_server>"
	transcfg := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	}

	client := http.Client{Transport:transcfg}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("request url failed\n",err)
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("read resp body failed")
	}
	fmt.Printf("html[:1024]\n %s", html[:1023])
}

