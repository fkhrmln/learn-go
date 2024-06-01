package gorouter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

type Logger struct {
	handler http.Handler
}

func (logger *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method, r.RequestURI)

	logger.handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	logger := Logger{
		handler: router,
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Hello World")
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &logger,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
