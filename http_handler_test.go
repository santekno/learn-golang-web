package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "mencetak Hello World",
			want: "Hello World\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/hello", nil)
			recorder := httptest.NewRecorder()
			HelloHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestSayHelloParameterHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return response with name",
			args: args{
				name: "santekno",
			},
			want: "Hello santekno",
		},
		{
			name: "success return response without name",
			args: args{
				name: "",
			},
			want: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s", tt.args.name), nil)
			recorder := httptest.NewRecorder()
			SayHelloParameterHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestMultipleParameterHandler(t *testing.T) {
	type args struct {
		firstName string
		lastName  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return response with name",
			args: args{
				firstName: "Santekno",
				lastName:  "Inc",
			},
			want: "Hello Santekno Inc",
		},
		{
			name: "success return response without name",
			args: args{
				firstName: "",
				lastName:  "",
			},
			want: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?first_name=%s&last_name=%s", tt.args.firstName, tt.args.lastName), nil)
			recorder := httptest.NewRecorder()
			MultipleParameterHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestMultipleParameterValueHandler(t *testing.T) {
	type args struct {
		name []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return response with name",
			args: args{
				name: []string{"santekno", "ihsan"},
			},
			want: "Hello santekno ihsan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s&name=%s", tt.args.name[0], tt.args.name[1]), nil)
			recorder := httptest.NewRecorder()
			MultipleParameterValueHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestRequestHedaerHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "set content type json",
			args: args{
				name: "santekno",
			},
			want: "santekno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/say", nil)
			request.Header.Add(X_POWERED_BY, tt.args.name)
			recorder := httptest.NewRecorder()
			RequestHedaerHandler(recorder, request)

			response := recorder.Result()
			poweredBy := response.Header.Get(X_POWERED_BY)

			if !reflect.DeepEqual(poweredBy, tt.want) {
				t.Errorf("poweredBy = %v, want %v", poweredBy, tt.want)
			}
		})
	}
}

func TestFormPostHandler(t *testing.T) {
	type args struct {
		firstName string
		lastName  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "set form post param",
			args: args{
				firstName: "ihsan",
				lastName:  "arif",
			},
			want: "ihsan arif",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody := strings.NewReader(fmt.Sprintf("first_name=%s&last_name=%s", tt.args.firstName, tt.args.lastName))
			request := httptest.NewRequest(http.MethodPost, "http://localhost/say", requestBody)
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			recorder := httptest.NewRecorder()
			FormPostHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("bodyString = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestResponseCodeHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
		wantCode int
	}{
		{
			name: "sent param name with value",
			args: args{
				name: "ihsan",
			},
			wantResp: "Hello ihsan",
			wantCode: http.StatusOK,
		},
		{
			name: "does't sent param name",
			args: args{
				name: "",
			},
			wantResp: "name is empty",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s", tt.args.name), nil)
			recorder := httptest.NewRecorder()
			ResponseCodeHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)
			code := response.StatusCode

			if !reflect.DeepEqual(bodyString, tt.wantResp) {
				t.Errorf("response = %v, want %v", bodyString, tt.wantResp)
			}

			if !reflect.DeepEqual(code, tt.wantCode) {
				t.Errorf("code = %v, want %v", code, tt.wantCode)
			}
		})
	}
}

func TestSetCookieHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "set cookie",
			args: args{
				name: "santekno",
			},
			want: "santekno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s", tt.args.name), nil)
			recorder := httptest.NewRecorder()
			SetCookieHandler(recorder, request)

			cookies := recorder.Result().Cookies()

			for _, cookie := range cookies {
				if !reflect.DeepEqual(cookie.Value, tt.want) {
					t.Errorf("response = %s, want %s", cookie.Value, tt.want)
				}
			}

		})
	}
}

func TestGetCookieHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get cookie handler without cookie",
			args: args{
				name: "",
			},
			want: "no cookie",
		},
		{
			name: "get cookie handler with cookie",
			args: args{
				name: "santekno",
			},
			want: "hello santekno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/say", nil)
			if tt.args.name != "" {
				cookie := new(http.Cookie)
				cookie.Name = "X-PXN-Name"
				cookie.Value = tt.args.name
				request.AddCookie(cookie)
			}

			recorder := httptest.NewRecorder()
			GetCookieHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %s, want %s", bodyString, tt.want)
			}
		})
	}
}

func TestServeFileWithEmbedHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get source test hello",
			args: args{
				name: "santekno",
			},
			want: `<html><h1>Test, Hello</h1></html>`,
		},
		{
			name: "resource not found",
			args: args{
				name: "",
			},
			want: `404 page not found`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/file?name=%s", tt.args.name), nil)
			recorder := httptest.NewRecorder()
			ServeFileWithEmbedHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestSimpleHTMLTemplateHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get from embed",
			want: "<html><body>Hello HTML Template</body></html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/file", nil)
			recorder := httptest.NewRecorder()
			SimpleHTMLTemplateHandler(recorder, request)

			body, _ := io.ReadAll(recorder.Result().Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %#v, want = %#v\n", bodyString, tt.want)
			}
		})
	}
}

func TestSimpleHTMLFileTemplateHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get from embed",
			want: "<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"UTF-8\">\n    <title>Hello Santekno, HTML File Template</title>\n  </head>\n  <body>\n    <h1>Hello Santekno, HTML File Template</h1>\n  </body>\n</html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/file", nil)
			recorder := httptest.NewRecorder()
			SimpleHTMLFileTemplateHandler(recorder, request)

			body, _ := io.ReadAll(recorder.Result().Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %#v, want = %#v\n", bodyString, tt.want)
			}
		})
	}
}

func TestTemplateEmbedHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get from embed",
			want: "<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"UTF-8\">\n    <title>Hello Santekno, HTML Embed Template</title>\n  </head>\n  <body>\n    <h1>Hello Santekno, HTML Embed Template</h1>\n  </body>\n</html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/file", nil)
			recorder := httptest.NewRecorder()
			TemplateEmbedHandler(recorder, request)

			body, _ := io.ReadAll(recorder.Result().Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %#v, want = %#v\n", bodyString, tt.want)
			}
		})
	}
}

func TestTemplateDataMapHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get from embed",
			want: "<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"UTF-8\">\n    <title>Template Data Map</title>\n  </head>\n  <body>\n    <h1>Hello Santekno</h1>\n  </body>\n</html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/file", nil)
			recorder := httptest.NewRecorder()
			TemplateDataMapHandler(recorder, request)

			body, _ := io.ReadAll(recorder.Result().Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %#v, want = %#v\n", bodyString, tt.want)
			}
		})
	}
}

func TestTemplateDataStructHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get from embed",
			want: "<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"UTF-8\">\n    <title>Template Data Struct</title>\n  </head>\n  <body>\n    <h1>Hello Santekno</h1>\n  </body>\n</html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/file", nil)
			recorder := httptest.NewRecorder()
			TemplateDataStructHandler(recorder, request)

			body, _ := io.ReadAll(recorder.Result().Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %#v, want = %#v\n", bodyString, tt.want)
			}
		})
	}
}
