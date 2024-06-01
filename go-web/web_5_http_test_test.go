package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPTest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	handler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
