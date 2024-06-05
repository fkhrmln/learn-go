package gofiber

import (
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestApp(t *testing.T) {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	err := app.Listen("127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
}
