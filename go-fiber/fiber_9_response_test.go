package gofiber

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestResponse(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := c.Query("name")

		return c.JSON(fiber.Map{
			"name": name,
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/?name=Fakhri%20Maulana%20Ihsan", nil)

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
