package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method, r.RequestURI)
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: http.HandlerFunc(handler),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
