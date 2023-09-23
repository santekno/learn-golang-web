package main

import (
	"html/template"
	"net/http"
)

func TemplateLayoutHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.html",
		"./templates/footer.html",
		"./templates/layout.html",
	))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Name":  "Santekno",
		"Title": "Contoh Template Layout",
	})
}
