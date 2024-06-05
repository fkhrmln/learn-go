package gorouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestNamedParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:productId/types/:typeId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productId := p.ByName("productId")
		typeId := p.ByName("typeId")

		fmt.Fprintf(w, "Product : %s\nType : %s\n", productId, typeId)
	})

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/products/1/types/1", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestCatchAllParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")

		fmt.Fprintln(w, "Image :", image)
	})

	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/images/pictures/go.png", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
