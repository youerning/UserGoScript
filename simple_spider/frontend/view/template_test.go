package view

import (
	"testing"
	"learngo/simple_spider/model"
	"os"
	frontendmodel "learngo/simple_spider/frontend/model"
)

func TestSearchResultView_Render(t *testing.T) {
	filename := "template.html"
	view := CreateSearchResultView(filename)


	outfile, err := os.Create("template_test.html")
	if err != nil {
		panic(err)
	}
	page := frontendmodel.SearchResult{}
	page.Hits = 123
	page.Start = 0
	item := model.Item{
		Url:"http://album.zhenai.com/u/106720604",
		Id:"106720604",
		Type:"zhenai",
		Payload: model.Profile{Name: "圆圆",
			Gender: "女", Age: 24, Height: 161, Weight: 56,
			Income: "3001-5000元", Marriage: "离异",
			Education: "高中及以下", Occupation: "美容师",
			HoKou: "吉林吉林", Xinzuo: "水瓶座",
			House: "租房", Car: "未购车"},
	}

	for i:=0;i<10;i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(outfile, page)
	if err != nil {
		panic(err)
	}
}