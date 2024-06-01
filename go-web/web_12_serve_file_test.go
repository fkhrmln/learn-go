package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/hello-world.html
var resourcesHelloWorldHTML string

//go:embed resources/not-found.html
var resourcesNotFoundHTML string

func TestServeFile(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name != "" {
			http.ServeFile(w, r, "resources/hello-world.html")
		} else {
			http.ServeFile(w, r, "resources/not-found.html")
		}
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

func TestServeFileEmbed(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name != "" {
			fmt.Fprintln(w, resourcesHelloWorldHTML)
		} else {
			fmt.Fprintln(w, resourcesNotFoundHTML)
		}
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
