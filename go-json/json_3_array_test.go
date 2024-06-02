package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeArray(t *testing.T) {
	type Address struct {
		Street, City, Country string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
		Hobbies []string
	}

	person := Person{
		Name:    "Fakhri Maulana Ihsan",
		Age:     20,
		Address: Address{"Mutiara Gading Timur", "Bekasi", "Indonesia"},
		Hobbies: []string{"Coding", "Gaming", "Sleeping"},
	}

	bytes, err := json.MarshalIndent(person, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	hobbies := []string{"Coding", "Gaming", "Sleeping"}

	bytes, err = json.MarshalIndent(hobbies, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestDecodeArray(t *testing.T) {
	type Address struct {
		Street, City, Country string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
		Hobbies []string
	}

	personJSON := `{
		"Name": "Fakhri Maulana Ihsan",
		"Age": 20,
		"Address": {
			"Street": "Mutiara Gading Timur",
			"City": "Bekasi",
			"Country": "Indonesia"
		},
		"Hobbies": ["Coding", "Gaming", "Sleeping"]
	}`

	person := Person{}

	json.Unmarshal([]byte(personJSON), &person)

	fmt.Println(person)

	hobbies := person.Hobbies

	json.Unmarshal([]byte(personJSON), &hobbies)

	fmt.Println(hobbies)
}
