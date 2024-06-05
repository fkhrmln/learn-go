package gofiber

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestErrorHandler(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusInternalServerError)

			return c.SendString("Error : " + err.Error())
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return errors.New("Internal Server Error")
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
