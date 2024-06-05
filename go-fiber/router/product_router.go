package router

import "github.com/gofiber/fiber/v2"

func ProductRouter(app *fiber.App) *fiber.Router {
	router := app.Group("/products")

	router.Get("/:productId", func(c *fiber.Ctx) error {
		productId := c.Params("productId")

		return c.SendString("Product : " + productId)
	})

	return &router
}
