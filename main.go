package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed resources
var resources embed.FS

func main() {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	mux.HandleFunc("/req", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	})

	mux.HandleFunc("/file", ServeFileHandler)
	mux.HandleFunc("/template", SimpleHTMLTemplateHandler)
	mux.HandleFunc("/template-file", SimpleHTMLFileTemplateHandler)
	mux.HandleFunc("/template-directory", TemplateDirectoryHanlder)
	mux.HandleFunc("/template-embed", TemplateEmbedHandler)

	mux.HandleFunc("/set-cookie", SetCookieHandler)
	mux.HandleFunc("/get-cookie", GetCookieHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
