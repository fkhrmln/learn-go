package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Powered-By", "Fakhri Maulana Ihsan")

		contentType := r.Header.Get("Content-Type")

		fmt.Fprintln(w, contentType)
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPost, "/", nil)

	request.Header.Add("Content-Type", "application/json")

	handler(recorder, request)

	response := recorder.Result()

	fmt.Println(response.Header.Get("X-Powered-By"))

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
