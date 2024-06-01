package goweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServer(t *testing.T) {
	dir := http.Dir("resources")

	fileServer := http.FileServer(dir)

	mux := http.ServeMux{}

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestFileServerEmbed(t *testing.T) {
	subDir, _ := fs.Sub(resources, "resources")

	fileServer := http.FileServer(http.FS(subDir))

	mux := http.ServeMux{}

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
