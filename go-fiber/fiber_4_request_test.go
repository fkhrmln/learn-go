package gofiber

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRequest(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-KEY")

		jwt := c.Cookies("JWT")

		return c.SendString(fmt.Sprintf("X-API-KEY : %s\nJWT : %s", apiKey, jwt))
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	request.Header.Add("X-API-KEY", "SECRET")

	request.AddCookie(&http.Cookie{
		Name:  "JWT",
		Value: "SECRET",
	})

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
