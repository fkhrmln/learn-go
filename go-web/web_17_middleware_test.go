package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

type Logger struct {
	handler http.Handler
}

func (logger *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method, r.RequestURI)

	logger.handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}

	logger := Logger{}

	logger.handler = http.HandlerFunc(handler)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &logger,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
