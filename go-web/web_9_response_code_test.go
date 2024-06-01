package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseCode(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			w.WriteHeader(http.StatusBadRequest)

			fmt.Fprintln(w, "Name is Empty")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Hello", name)
		}
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPost, "/?name=Fakhri%20Maulana%20Ihsan", nil)

	handler(recorder, request)

	response := recorder.Result()

	fmt.Println(response.Status)

	fmt.Println(response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
