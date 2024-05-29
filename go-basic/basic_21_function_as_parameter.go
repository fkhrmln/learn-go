package main

import (
	"fmt"
	"strings"
)

type ToUppercase func(string) string

func sayHelloUppercase(name string, toUppercase ToUppercase) {
	fmt.Println("HELLO", toUppercase(name))
}

// func sayHelloUppercase(name string, toUppercase func(string) string) {
// 	fmt.Println("HELLO", toUppercase(name))
// }

func toUppercase(value string) string {
	return strings.ToUpper(value)
}

func main() {
	sayHelloUppercase("Fakhri", toUppercase)
}
