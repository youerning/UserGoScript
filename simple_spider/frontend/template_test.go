package frontend

import (
	"testing"
	"html/template"
)

func TestTemplate(t *testing.T) {
	template := template.Must(
		template.ParseFiles("template.html"))
}