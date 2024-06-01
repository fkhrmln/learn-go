package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func redirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Dashboard")
}

func redirectFrom(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "admin" {
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
	} else {
		fmt.Fprintln(w, "Hello World")
	}
}

func TestRedirect(t *testing.T) {
	mux := http.ServeMux{}

	mux.HandleFunc("/", redirectFrom)

	mux.HandleFunc("/dashboard", redirectTo)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
