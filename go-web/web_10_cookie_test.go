package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{}
	cookie.Name = "JWT"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, &cookie)

	fmt.Fprintln(w, "Cookie Created")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("JWT")

	if err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprintln(w, "Cookie Not Found")
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Hello", cookie.Value)
}

func TestCookie(t *testing.T) {
	mux := http.ServeMux{}

	mux.HandleFunc("/", GetCookie)

	mux.HandleFunc("/set-cookie", SetCookie)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/?name=Fakhri%20Maulana%20Ihsan", nil)

	SetCookie(recorder, request)

	response := recorder.Result()

	for _, cookie := range response.Cookies() {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestGetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cookie := http.Cookie{}
	cookie.Name = "JWT"
	cookie.Value = "Fakhri Maulana Ihsan"
	cookie.Path = "/"

	request.AddCookie(&cookie)

	GetCookie(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
