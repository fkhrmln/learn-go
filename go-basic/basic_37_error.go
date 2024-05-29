package main

import (
	"errors"
	"fmt"
)

func connectDatabase(value bool) (string, error) {
	if value == false {
		return "", errors.New("Database Failed")
	}

	return "Database Connected", nil
}

type ValidationError struct {
	Message string
}

func (validationError *ValidationError) Error() string {
	return validationError.Message
}

type NotFoundError struct {
	Message string
}

func (notFoundError *NotFoundError) Error() string {
	return notFoundError.Message
}

func findName(value any, names ...string) (string, error) {
	var isFound bool

	switch value.(type) {
	case int:
		return "", &ValidationError{"Can't Find Number"}
	case bool:
		return "", &ValidationError{"Can't Find Bool"}
	}

	for _, name := range names {
		if name == value {
			isFound = true

			break
		} else {
			isFound = false
		}
	}

	if isFound == false {
		return "", &NotFoundError{value.(string) + " Not Found"}
	}

	return value.(string) + " Found", nil
}

func main() {
	result, err := connectDatabase(false)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(result)
	}

	names := []string{"Fakhri", "Rifky", "Audry"}

	isFound, err := findName("Bobby", names...)

	if err != nil {
		switch errorType := err.(type) {
		case *ValidationError:
			fmt.Println("Validation Error :", errorType)
		case *NotFoundError:
			fmt.Println("Not Found Error :", errorType)
		default:
			fmt.Println("Error : Unknown Error")
		}
	} else {
		fmt.Println(isFound)
	}
}
