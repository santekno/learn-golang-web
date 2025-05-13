package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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
