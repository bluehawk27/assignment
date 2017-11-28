package httpapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	handler.ServeHTTP(w, req)

	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	a := args{
		w:   w,
		req: req,
	}
	tests := []struct {
		name string
		args args
	}{
		{"HTTP-PING", a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ping(tt.args.w, tt.args.req)
		})
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func TestAdd(t *testing.T) {
	req, err := http.NewRequest("POST", "/add/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = ioutil.NopCloser(bytes.NewReader([]byte("bar")))

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(Add)

	handler.ServeHTTP(w, req)

	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	a := args{
		w:   w,
		req: req,
	}
	tests := []struct {
		name string
		args args
	}{
		{"HTTP-ADD", a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.w, tt.args.req)
		})
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func TestGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/get/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)

	handler.ServeHTTP(w, req)
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}

	a := args{
		w:   w,
		req: req,
	}
	tests := []struct {
		name string
		args args
	}{
		{"HTTP-Get", a},
	}
	for _, tt := range tests {
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		t.Run(tt.name, func(t *testing.T) {
			Get(tt.args.w, tt.args.req)
		})
	}
}
