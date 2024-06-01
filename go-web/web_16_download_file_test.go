package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Query().Get("file")

		if file == "" {
			w.WriteHeader(http.StatusBadRequest)

			fmt.Fprintln(w, "Bad Request")

			return
		}

		w.WriteHeader(http.StatusOK)

		w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")

		http.ServeFile(w, r, "images/"+file)
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
