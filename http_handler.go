package main

import (
	"fmt"
	"net/http"
	"strings"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func SayHelloParameterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func MultipleParameterHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	if firstName == "" && lastName == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
	}
}

func MultipleParameterValueHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	if len(names) == 0 {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
	}
}

const X_POWERED_BY = "X-Powered-By"

func RequestHedaerHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "%s %s", firstName, lastName)
}

func ResponseCodeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PXN-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Success create cookie")
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PXN-Name")
	if err != nil {
		fmt.Fprint(w, "no cookie")
	} else {
		fmt.Fprintf(w, "hello %s", cookie.Value)
	}
}
