package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserCity(t *testing.T) {
	content, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(content))
	result := ParserCity(content)
	const userSize  = 44
	if len(result.Requests) != userSize {
		t.Errorf("parser error: user size should be %d, but got %d", userSize, len(result.Requests))
	}
}