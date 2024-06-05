package gofiber

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

//go:embed source/go.png
var image []byte

func TestUploadFile(t *testing.T) {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")

		if err != nil {
			panic(err)
		}

		err = c.SaveFile(file, "target/"+file.Filename)

		if err != nil {
			panic(err)
		}

		return c.SendString("Upload File Success")
	})

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	file, err := writer.CreateFormFile("file", "go.png")

	if err != nil {
		panic(err)
	}

	file.Write(image)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/", body)

	request.Header.Add("Content-Type", writer.FormDataContentType())

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
