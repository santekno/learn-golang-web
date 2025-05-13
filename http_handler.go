package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// logic web
	fmt.Fprint(w, "hello world")
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi")
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method)
	fmt.Fprint(w, r.RequestURI)
}
