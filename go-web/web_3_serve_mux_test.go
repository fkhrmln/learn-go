package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.ServeMux{}

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi World")
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Products")
	})

	mux.HandleFunc("/products/categories/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Categories")
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
