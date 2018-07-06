package view

import (
	"html/template"
	"io"
	"learngo/simple_spider/frontend/model"
)

type SearchResultView struct {
	Tpl *template.Template
}

func CreateSearchResultView(filename string) SearchResultView{
	return SearchResultView{
		Tpl: template.Must(
			template.ParseFiles(filename),
		),
	}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error{
	return s.Tpl.Execute(w, data)
}
