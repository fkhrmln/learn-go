package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestQueryParameter(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprintln(w, "Hello Anonymous")
		} else {
			fmt.Fprintln(w, "Hello", name)
		}
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/?name=Fakhri%20Maulana%20Ihsan", nil)

	handler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleQueryParameter(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		fmt.Fprintln(w, strings.Join(query["name"], " "))
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/?name=Fakhri&name=Rifky&name=Audry", nil)

	handler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
