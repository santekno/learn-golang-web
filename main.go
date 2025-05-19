package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed resources
var resources embed.FS

func main() {
	// directory := http.Dir("./resources")
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/hi", HiHandler)
	mux.HandleFunc("/request", RequestHandler)
	mux.HandleFunc("/sayhi", SayHalloParameterHandler)
	mux.HandleFunc("/sayhi-multiple", MultipleParameterHandler)
	mux.HandleFunc("/sayhi-muletiplevalue", MultipleParameterValueHandler)
	mux.HandleFunc("/set-header", RequestHeaderHandler)
	mux.HandleFunc("/response-code", ResponseCodeHandler)
	mux.HandleFunc("/set-cookie", SetCookieHandler)
	mux.HandleFunc("/get-cookie", GetCookieHandler)

	// static file from resources folder
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
