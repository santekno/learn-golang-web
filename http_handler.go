package main

import (
	"fmt"
	"html/template"
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
	cookie.Name = "X-Santekno-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "success create cookie")
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Santekno-Name")
	if err != nil {
		fmt.Fprint(w, "no cookie")
	} else {
		fmt.Fprintf(w, "hello %s", cookie.Value)
	}
}

func SimpleHTMLTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{ . }}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func SimpleHTMLFileTemplateHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/simple.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.html", "Hello santekno, HTML File Template")
}

func TemplateDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*html"))
	t.ExecuteTemplate(w, "simple.html", "Hello santekno, HTML directory file template")
}

// // go:embed templates/*.html
// var templates embed.FS

// func TemplateEmbedHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFS(templates, "templates/*.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	t.ExecuteTemplate(w, "simple.html", "Hello santekno, HTML embed template")
// }

func TemplateDataMapHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Santekno",
	})
}

type Page struct {
	Title string
	Name  string
}

func TemplateDataStructHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(w, "name.html", Page{
		Title: "Template Data Struct",
		Name:  "Santekno",
	})
}

func TemplateActionIfHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.html"))
	t.ExecuteTemplate(w, "if.html", map[string]interface{}{
		"Name": "Santekno",
	})
}

func TemplateActionComparatorHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.html"))
	t.ExecuteTemplate(w, "comparator.html", map[string]interface{}{
		"Name":  "Santekno",
		"Value": 70,
	})
}

func TemplateActionRangeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.html"))
	t.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Hobbies": []string{
			"Gaming", "Badminton", "Coding", "Reading",
		},
	})
}

func TemplateActionWithHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.html"))
	t.ExecuteTemplate(w, "with.html", map[string]interface{}{
		"Name": "Santekno",
		"Address": map[string]interface{}{
			"Street": "Jalan Padjadjaran",
			"City":   "Bogor",
		},
	})
}
