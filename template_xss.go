package main

import (
	"html/template"
	"net/http"
)

func TemplateAutoEscapeHandler(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]interface{}{
		"Title": "Golang Tutorial Santekno Auto Escape",
		"Body":  "<p>Selamat Belajar Golang Auto Escape Santekno<script>alert('Halo Anda ke hack')</script></p>",
	})
}

func TemplateDisabledAutoEscapeHandler(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]interface{}{
		"Title": "Golang Tutorial Santekno Auto Escape",
		"Body":  template.HTML("<p>Selamat Belajar Golang Auto Escape Santekno<script>alert('Halo Anda ke hack')</script></p>"),
	})
}

func TemplateXSSAttackHandler(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.html", map[string]interface{}{
		"Title": "Golang Tutorial Santekno Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}
