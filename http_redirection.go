package main

import (
	"fmt"
	"net/http"
)

func RedirectToHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Santekno")
}

func RedirectFromHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}
