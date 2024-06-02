package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTagEncode(t *testing.T) {
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

	bytes, err := json.MarshalIndent(product, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestTagDecode(t *testing.T) {
	type Product struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
		Image string `json:"image"`
	}

	productJSON := `{
		"id": "00001",
		"name": "Product 1",
		"price": 1000000,
		"image": "https://images/product.png"
	}`

	product := Product{}

	err := json.Unmarshal([]byte(productJSON), &product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
