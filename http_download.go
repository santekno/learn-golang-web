package main

import (
	"fmt"
	"net/http"
)

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
	}

	w.Header().Add("Content-Disposition", "attachment;filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./resources/"+fileName)
}
