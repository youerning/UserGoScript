package controller

import (
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"learngo/simple_spider/frontend/view"
	"context"
	"learngo/simple_spider/frontend/model"
	"reflect"
	model2 "learngo/simple_spider/model"
)

type FrontEndHandler struct {
	view view.SearchResultView
	client *elastic.Client
}


func CreateFrontendHandler(template string) FrontEndHandler {
	client , err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://192.168.110.104:9200"),
	)
	if err != nil {
		panic(err)
	}
	v := view.CreateSearchResultView(template)

	return FrontEndHandler{
		view: v,
		client:client,
	}
}

// handle data from localhost/search?q=男 已购房&from=20
func (h FrontEndHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}


	var page model.SearchResult

	ret, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}


	page.Start = from
	page.Hits = ret.TotalHits()
	//page.Items = ret.

	//fmt.Fprintf(w, "q=%s from=%d", q, from)
	page.Items = ret.Each(reflect.TypeOf(model2.Item{}))
	//fmt.Printf("%+v\n",page)
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
