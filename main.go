package main

import (
	"net/http"
)

func main() {
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

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
