package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.html
var templatesHTML embed.FS

func TestTemplate(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/hello-world.html")

		if err != nil {
			panic(err)
		}

		name := r.URL.Query().Get("name")

		data := map[string]interface{}{
			"Title": "Hello World",
			"Name":  name,
		}

		if name != "" {
			t.ExecuteTemplate(w, "hello-world.html", data)
		} else {
			data["Name"] = "Hello Anonymous"

			t.ExecuteTemplate(w, "hello-world.html", data)
		}
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/?name=Fakhri%20Maulana%20Ihsan", nil)

	handler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestTemplateEmbed(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(templatesHTML, "templates/*.html")

		if err != nil {
			panic(err)
		}

		name := r.URL.Query().Get("name")

		data := map[string]interface{}{
			"Title": "Hello World",
			"Name":  name,
		}

		if name != "" {
			t.ExecuteTemplate(w, "hello-world.html", data)
		} else {
			data["Name"] = "Hello Anonymous"

			t.ExecuteTemplate(w, "hello-world.html", data)
		}
	}

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/?name=Fakhri%20Maulana%20Ihsan", nil)

	handler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
