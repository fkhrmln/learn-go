package gofiber

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestParams(t *testing.T) {
	app := fiber.New()

	app.Get("/products/:productId/types/:typeId", func(c *fiber.Ctx) error {
		productId := c.Params("productId")

		typeId := c.Params("typeId")

		return c.SendString(fmt.Sprintf("Product : %s\nType : %s", productId, typeId))
	})

	request := httptest.NewRequest(http.MethodGet, "/products/1/types/1", nil)

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
