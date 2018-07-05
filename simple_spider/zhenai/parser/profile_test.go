package parser

import (
	"io/ioutil"
	"learngo/simple_spider/model"
	"testing"
)

func TestParserProfile(t *testing.T) {
	content, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(content))
	result := ParserProfile(content, "http://album.zhenai.com/u/106720604", "圆圆")

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

	if item != result.Items[0] {
		t.Errorf("parser error: result size should be %#v, but got %#v", item, result.Items[0])
	}
}
