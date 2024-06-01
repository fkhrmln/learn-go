package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostForm(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		age := r.PostFormValue("age")
		address := r.PostFormValue("address")

		fmt.Fprintln(w, name, age, address)
	}

	recorder := httptest.NewRecorder()

	form := strings.NewReader("name=Fakhri Maulana Ihsan&age=20&address=Mutiara Gading Timur")

	request := httptest.NewRequest(http.MethodPost, "/", form)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	handler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
