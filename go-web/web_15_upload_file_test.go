package goweb

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestUploadFile(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")

		if err != nil {
			panic(err)
		}

		unixTime := strconv.FormatInt(time.Now().Unix(), 10)

		fileDestination, err := os.Create(fmt.Sprintf("images/%s.%s", unixTime, strings.Split(fileHeader.Filename, ".")[1]))

		if err != nil {
			panic(err)
		}

		_, err = io.Copy(fileDestination, file)

		if err != nil {
			panic(err)
		}

		fmt.Fprintln(w, "Upload File Success")
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
