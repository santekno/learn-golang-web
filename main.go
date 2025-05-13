package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hi")
	})

	mux.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
