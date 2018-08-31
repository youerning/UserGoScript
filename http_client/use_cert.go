// use specified certificate signed by unknown authority for request
package main

import (
	"crypto/tls"
	"log"
	"io/ioutil"
	"crypto/x509"
	"net/http"
	"fmt"
)

func main() {
	url  := "<https_server>"
	certFile := "<path_to_cert>"
	keyFile := "<path_to_key>"

	// load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)

	if err != nil {
		log.Fatal(err)
	}

	// load CA cert
	caCert, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	//	setup http client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: caCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport:transport}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}
