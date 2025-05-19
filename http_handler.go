package main

import (
	"fmt"
	"net/http"
	"strings"
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

func SayHalloParameterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %s", name)
	}
}

func MultipleParameterHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	if firstName == "" && lastName == "" {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %s %s", firstName, lastName)
	}
}

func MultipleParameterValueHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	if len(names) == 0 {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %s", strings.Join(names, " "))
	}
}

const X_POWERED_BY = "X-Powered-By"

func RequestHeaderHandler(w http.ResponseWriter, r *http.Request) {
	poweredBy := r.Header.Get(X_POWERED_BY)
	w.Header().Add(X_POWERED_BY, poweredBy)
	fmt.Fprint(w, poweredBy)
}

func FormPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")
	fmt.Fprintf(w, "first_name: %s last_name: %s", firstName, lastName)
}
