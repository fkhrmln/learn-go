package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func findName(value interface{}, names ...string) error {
	var isFound bool

	if _, ok := value.(string); !ok {
		return ValidationError
	}

	for _, name := range names {
		if name == value {
			isFound = true

			break
		}

		isFound = false
	}

	if isFound == false {
		return NotFoundError
	}

	return nil
}

func main() {
	names := []string{"Fakhri", "Rifky", "Audry"}

	err := findName("Bobby", names...)

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Validation Error")
		} else {
			fmt.Println("Unknown Error")
		}
	} else {
		fmt.Println("Found")
	}
}
