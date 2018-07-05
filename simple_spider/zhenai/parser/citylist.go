package parser

import (
	"learngo/simple_spider/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]+>([^<]+)</a>`

// nil parser for placehold
func NilParser(content []byte) engine.ParserResult {
	return engine.ParserResult{}
}

func ParserCityList(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	//var limit = 10
	for _, v := range match {
		//result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: ParserCity,
		})
		//limit--
		//if limit <0 {
		//	break
		//}
	}
	return result
}






