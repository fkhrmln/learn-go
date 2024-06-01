package gorouter

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

//go:embed images
var images embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()

	subDir, _ := fs.Sub(images, "images")

	router.ServeFiles("/images/*filepath", http.FS(subDir))

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
