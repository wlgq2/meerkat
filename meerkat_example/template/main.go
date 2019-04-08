package main

import (
	"github.com/wlgq2/meerkat"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (template *Template) Render(writer io.Writer, name string, data interface{}, context *meerkat.Context) error {
	return template.templates.ExecuteTemplate(writer, name, data)
}

func main() {
	server := meerkat.New()
	server.SetRender(&Template{templates: template.Must(template.ParseGlob("*.html")),})
	server.GET("/hello", func (context *meerkat.Context) error {
		return context.Render(http.StatusOK, "hello", "World")
	})
	meerkat.LogInstance().Fatalln(server.Start(":8001"))
}
