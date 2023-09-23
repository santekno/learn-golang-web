package main

import (
	"html/template"
	"net/http"
)

func TemplateDataMapHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Santekno",
	})
}

func TemplateDataStructHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", Page{
		Title: "Template Data Struct",
		Name:  "Santekno",
	})
}
