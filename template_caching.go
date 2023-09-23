package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, `templates/*.html`))

func TemplateCachingHandler(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.html", "Hello HTML Template")
}
