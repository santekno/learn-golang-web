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
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

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
	// mux.HandleFunc("/template-embed", TemplateEmbedHandler)
	mux.HandleFunc("/template-data-map", TemplateDataMapHandler)
	mux.HandleFunc("/template-data-struct", TemplateDataStructHandler)
	mux.HandleFunc("/template-action-if", TemplateActionIfHandler)
	mux.HandleFunc("/template-action-comparator", TemplateActionComparatorHandler)
	mux.HandleFunc("/template-action-range", TemplateActionRangeHandler)
	mux.HandleFunc("/template-action-with", TemplateActionWithHandler)
	mux.HandleFunc("/template-layout", TemplateLayoutHandler)
	mux.HandleFunc("/template-function", TemplateFunctionHandler)
	mux.HandleFunc("/template-global-function", TemplateGlobalFunctionHandler)
	mux.HandleFunc("/template-manual-global-function", TemplateManualGlobalFunctionHandler)
	mux.HandleFunc("/template-function-pipelines", TemplateFunctionPipelinesHandler)
	mux.HandleFunc("/template-cache", TemplateCachingHandler)
	mux.HandleFunc("/template-auto-escape", TemplateAutoEscapeHandler)
	mux.HandleFunc("/template-disable-auto-escape", TemplateDisabledAutoEscapeHandler)
	mux.HandleFunc("/template-xss-attack", TemplateXSSAttackHandler)

	mux.HandleFunc("/redirect-to", RedirectToHandler)
	mux.HandleFunc("/redirect-from", RedirectFromHandler)

	mux.HandleFunc("/upload-form", UploadFormHandler)
	mux.HandleFunc("/upload", UploadHandler)

	mux.HandleFunc("/download", DownloadFileHandler)

	mux.HandleFunc("/set-cookie", SetCookieHandler)
	mux.HandleFunc("/get-cookie", GetCookieHandler)

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("upps")
	})

	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux

	errMiddleware := &ErrorMiddleware{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
