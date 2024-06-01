package gorouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:productId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("productId")

		fmt.Fprintln(w, "Product :", productId)
	})

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/products/1", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
