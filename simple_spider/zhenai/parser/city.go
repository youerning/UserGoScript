package parser

import (
	"learngo/simple_spider/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityRe    = regexp.MustCompile(`<a[^h]+href="(http://www.zhenai.com/zhenghun/[^"]+)">[^<]+</a>`)
)

func ParserCity(content []byte) engine.ParserResult {
	match := profileRe.FindAllSubmatch(content, -1)
	//fmt.Printf("match: %s\n", match)

	result := engine.ParserResult{}
	for _, v := range match {
		name := string(v[2])
		url := string(v[1])
		//result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParserResult {
				return ParserProfile(content, url, name)
			},
		})
	}

	matchs := cityRe.FindAllSubmatch(content, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}
	return result
}
