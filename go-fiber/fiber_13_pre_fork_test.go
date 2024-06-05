package gofiber

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestPreFork(t *testing.T) {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Prefork:      true,
	})

	if fiber.IsChild() {
		fmt.Println("Child Process")
	} else {
		fmt.Println("Parent Process")
	}

	err := app.Listen("127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
}
