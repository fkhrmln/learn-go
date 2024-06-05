package gofiber

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestForm(t *testing.T) {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		name := c.FormValue("name")

		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Fakhri Maulana Ihsan")

	request := httptest.NewRequest(http.MethodPost, "/", body)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
