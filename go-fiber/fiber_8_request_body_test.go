package gofiber

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRequestBody(t *testing.T) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		loginRequest := &LoginRequest{}

		err := c.BodyParser(loginRequest)

		if err != nil {
			panic(err)
		}

		return c.SendString("Hello " + loginRequest.Username)
	})

	loginRequest := &LoginRequest{
		Username: "Fakhri Maulana Ihsan",
		Password: "SECRET",
	}

	bytes, err := json.MarshalIndent(loginRequest, "", "    ")

	if err != nil {
		panic(err)
	}

	body := strings.NewReader(string(bytes))

	request := httptest.NewRequest(http.MethodPost, "/", body)

	request.Header.Add("Content-Type", "application/json")

	response, err := app.Test(request)

	if err != nil {
		panic(err)
	}

	bytes, err = io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
