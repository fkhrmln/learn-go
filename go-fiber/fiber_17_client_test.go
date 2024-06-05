package gofiber

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestClient(t *testing.T) {
	client := fiber.AcquireClient()

	defer fiber.ReleaseClient(client)

	agent := client.Get("https://jsonplaceholder.typicode.com/users")

	agent.JSON(fiber.Map{
		"username": "Fakhri Maulana Ihsan",
		"password": "SECRET",
	})

	agent.ContentType("application/json")

	_, response, errors := agent.String()

	if errors != nil {
		panic(errors)
	}

	fmt.Println(response)
}
