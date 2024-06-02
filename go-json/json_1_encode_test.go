package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	type Person struct {
		Name    string
		Age     int
		Address string
	}

	persons := []Person{
		{
			Name:    "Fakhri Maulana Ihsan",
			Age:     20,
			Address: "Mutiara Gading Timur",
		},
		{
			Name:    "Rifky Ferdiansyah",
			Age:     20,
			Address: "Bekasi Timur Regency",
		},
		{
			Name:    "Audry Elgalia",
			Age:     20,
			Address: "Rawa Kalong",
		},
	}

	bytes, err := json.MarshalIndent(persons, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
