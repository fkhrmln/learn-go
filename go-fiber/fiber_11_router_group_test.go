package gofiber

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRouterGroup(t *testing.T) {
	app := fiber.New()

	userRouter := app.Group("/users")

	userRouter.Get("/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")

		return c.SendString("User : " + userId)
	})

	productRouter := app.Group("/products")

	productRouter.Get("/:productId", func(c *fiber.Ctx) error {
		productId := c.Params("productId")

		return c.SendString("Product : " + productId)
	})

	request := httptest.NewRequest(http.MethodGet, "/users/1", nil)

	response, err := app.Test(request)

	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(response.Body)

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
