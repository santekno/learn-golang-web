package main

import (
	"fmt"
	"net/http"
)

type ErrorMiddleware struct {
	Handler http.Handler
}

func (middleware *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("recover :", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
	}()
	middleware.Handler.ServeHTTP(w, r)
}
