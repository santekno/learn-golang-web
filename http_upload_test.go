package main

import (
	"bytes"
	_ "embed"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//go:embed resources/tutorial-golang.webp
var uploadFileTest []byte

func TestUploadHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				name: "success upload",
			},
			want: "<!DOCTYPE html>\n<html>\n  <head>\n    <title>Success success upload</title>\n  </head>\n  <body>\n    <h1>success upload</h1>\n    <a href=\"/static/contoh-upload.jpg\">File</a>\n  </body>\n</html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := new(bytes.Buffer)

			writer := multipart.NewWriter(body)
			writer.WriteField("name", tt.args.name)
			file, _ := writer.CreateFormFile("file", "contoh-upload.jpg")
			file.Write(uploadFileTest)
			writer.Close()

			request := httptest.NewRequest(http.MethodPost, "http://localhost/upload", body)
			request.Header.Set("Content-Type", writer.FormDataContentType())
			recorder := httptest.NewRecorder()
			UploadHandler(recorder, request)

			bodyResponse, _ := io.ReadAll(recorder.Result().Body)
			bodyString := string(bodyResponse)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %#v, want = %#v\n", bodyString, tt.want)
			}
		})
	}
}
