package router

import "github.com/gofiber/fiber/v2"

func UserRouter(app *fiber.App) *fiber.Router {
	router := app.Group("/users")

	router.Get("/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")

		return c.SendString("User : " + userId)
	})

	return &router
}
