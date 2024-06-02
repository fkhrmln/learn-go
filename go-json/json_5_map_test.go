package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "00001",
		"name":  "Product 1",
		"price": 1000000,
		"image": "https://images/product.png",
	}

	bytes, err := json.MarshalIndent(product, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestMapDecode(t *testing.T) {
	productJSON := `{
		"id": "00001",
		"name": "Product 1",
		"price": 1000000,
		"image": "https://images/product.png"
	}`

	product := map[string]interface{}{}

	err := json.Unmarshal([]byte(productJSON), &product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
