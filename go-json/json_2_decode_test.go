package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	type Person struct {
		Name    string
		Age     int
		Address string
	}

	personJSON := `{
		"Name": "Fakhri Maulana Ihsan",
		"Age": 20,
		"Address": "Mutiara Gading Timur"
	}`

	person := Person{}

	json.Unmarshal([]byte(personJSON), &person)

	fmt.Println(person)
}
