package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	type Product struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Image string `json:"image"`
	}

	product := Product{
		Id:    "00001",
		Name:  "Product 1",
		Price: 1000000,
		Image: "https://images/product.png",
	}

	writer, err := os.Create("product.json")

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer)

	encoder.SetIndent("", "    ")

	err = encoder.Encode(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

func TestStreamingDecoder(t *testing.T) {
	type Product struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Image string `json:"image"`
	}

	product := Product{}

	reader, err := os.Open("product.json")

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
