package persist

import (
	"learngo/simple_spider/model"
	"testing"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	expected := model.Item{
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


	// TODO: Try to start up elasticsearch in docker
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.110.104:9200"),
	)

	const index = "dating_profile_test"
	err = save(client, expected, index)
	if err != nil {
		t.Errorf("save error")
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id( expected.Id).
		Do(context.Background())
	if err != nil {
		t.Errorf("get item: %v \nerror: ", expected.Id, err)
	}

	var actual model.Item
	err = json.Unmarshal(
		*resp.Source, &actual)

	actual.Payload = model.FromJson(actual.Payload)
	if err != nil {
		t.Errorf("unmarshal error: %v", err)
	}

	//t.Logf("Got %+v\nexpected %+v", actual, expected)
	if actual != expected{
		t.Errorf("Got %+v, expected %+v", actual, expected)
	}

}
