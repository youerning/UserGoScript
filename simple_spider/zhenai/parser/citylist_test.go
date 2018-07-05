package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserCitylList(t *testing.T) {
	
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(content))
	result := ParserCityList(content)
	const citySize  = 470
	if len(result.Requests) != citySize {
		t.Errorf("parser error: result size should be %d, but got %d", citySize, len(result.Requests))
	}
}