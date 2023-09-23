package main

import (
	"html/template"
	"net/http"
	"strings"
)

func TemplateFunctionHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").
		Parse(`{{ .SayHello "Santekno" }}`))

	t.ExecuteTemplate(w, "FUNCTION", Page{
		Title: "Hello",
		Name:  "Santekno",
	})
}

func TemplateGlobalFunctionHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").
		Parse(`{{ len .Name }}`))

	t.ExecuteTemplate(w, "FUNCTION", Page{
		Title: "Hello",
		Name:  "Santekno",
	})
}

func TemplateManualGlobalFunctionHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(template.FuncMap{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))
	t.ExecuteTemplate(w, "FUNCTION", Page{
		Title: "Hello",
		Name:  "Santekno",
	})
}

func TemplateFunctionPipelinesHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(template.FuncMap{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))
	t.ExecuteTemplate(w, "FUNCTION", Page{
		Title: "Hello",
		Name:  "Santekno",
	})
}
