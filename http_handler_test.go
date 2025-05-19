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
			name: "print hello world",
			want: "hello world",
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

func TestHiHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "print hello world",
			want: "hi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/hello", nil)
			recorder := httptest.NewRecorder()
			HiHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestRequestHandler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "print hello world",
			want: "GEThttp://localhost/request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/request", nil)
			recorder := httptest.NewRecorder()
			RequestHandler(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)
			bodyString := string(body)

			if !reflect.DeepEqual(bodyString, tt.want) {
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}

func TestSayHalloParameterHandler(t *testing.T) {
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
			want: "hello santekno",
		},
		{
			name: "success return response without name",
			args: args{
				name: "",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s", tt.args.name), nil)
			recorder := httptest.NewRecorder()
			SayHalloParameterHandler(recorder, request)

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
				firstName: "ihsan",
				lastName:  "arif",
			},
			want: "hello ihsan arif",
		},
		{
			name: "success return response without name",
			args: args{
				firstName: "",
				lastName:  "",
			},
			want: "hello",
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
				name: []string{"ihsan", "arif"},
			},
			want: "hello ihsan arif",
		},
		{
			name: "success return response without name",
			args: args{
				name: []string{},
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var request *http.Request
			if len(tt.args.name) != 0 {
				request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost/say?name=%s&name=%s", tt.args.name[0], tt.args.name[1]), nil)
			} else {
				request = httptest.NewRequest(http.MethodGet, "http://localhost/say", nil)
			}

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

func TestRequestHeaderHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return with header",
			args: args{
				name: "santekno",
			},
			want: "santekno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "http://localhost/say?", nil)
			recorder := httptest.NewRecorder()
			request.Header.Add(X_POWERED_BY, tt.args.name)

			RequestHeaderHandler(recorder, request)

			response := recorder.Result()
			poweredBy := response.Header.Get(X_POWERED_BY)

			if !reflect.DeepEqual(poweredBy, tt.want) {
				t.Errorf("response = %v, want %v", poweredBy, tt.want)
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
			want: "first_name: ihsan last_name: arif",
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
				t.Errorf("response = %v, want %v", bodyString, tt.want)
			}
		})
	}
}
