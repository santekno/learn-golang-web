package main

import (
	"html/template"
	"net/http"
)

func TemplateActionIfHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.html"))
	t.ExecuteTemplate(w, "if.html", map[string]interface{}{
		"Name": "Santekno",
	})
}

func TemplateActionComparatorHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.html"))
	t.ExecuteTemplate(w, "comparator.html", map[string]interface{}{
		"Name":  "Santekno",
		"Value": 60,
	})
}

func TemplateActionRangeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.html"))
	t.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Hobbies": []string{
			"Gaming", "Reading", "Coding",
		},
	})
}

func TemplateActionWithHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.html"))
	t.ExecuteTemplate(w, "address.html", map[string]interface{}{
		"Name": "Santekno",
		"Address": map[string]interface{}{
			"Street": "Jalan Padjadjaran",
			"City":   "Bogor",
		},
	})
}
