package main

import (
	"io"
	"net/http"
	"os"
)

func UploadFormHandler(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.html", nil)
	if err != nil {
		panic(err)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}
