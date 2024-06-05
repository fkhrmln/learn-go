package gofiber

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestDownloadFile(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		file := "source/hello-world.txt"

		return c.Download(file)
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	response, err := app.Test(request)

	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
