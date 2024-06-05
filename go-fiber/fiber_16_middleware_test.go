package gofiber

import (
	"fmt"
	"go-fiber/router"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Method(), c.Path())

		err := c.Next()

		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	router.UserRouter(app)

	router.ProductRouter(app)

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

	request = httptest.NewRequest(http.MethodGet, "/users/1", nil)

	response, err = app.Test(request)

	if err != nil {
		panic(err)
	}

	bytes, err = io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	request = httptest.NewRequest(http.MethodGet, "/products/1", nil)

	response, err = app.Test(request)

	if err != nil {
		panic(err)
	}

	bytes, err = io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
