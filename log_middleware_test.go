package main

import (
	"net/http"
	"testing"
)

func TestLogMiddleware_ServeHTTP(t *testing.T) {
	type fields struct {
		Handler http.Handler
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := &LogMiddleware{
				Handler: tt.fields.Handler,
			}
			middleware.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
