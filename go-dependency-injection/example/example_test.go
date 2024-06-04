package example

import (
	"fmt"
	"testing"
)

func TestDependencyInjectionSuccess(t *testing.T) {
	exampleController, err := InitializeExampleController(false)

	if err != nil {
		panic(err)
	}

	fmt.Println(exampleController)
}

func TestDependencyInjectionError(t *testing.T) {
	exampleController, err := InitializeExampleController(true)

	if err != nil {
		panic(err)
	}

	fmt.Println(exampleController)
}
