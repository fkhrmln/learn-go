package main

import "fmt"

func main() {
	name := "Fakhri"

	switch name {
	case "Fakhri":
		fmt.Println("Hello Fakhri")
	case "Rifky":
		fmt.Println("Hello Rifky")
	case "Audry":
		fmt.Println("Hello Audry")
	default:
		fmt.Println("Hello", name)
	}

	password := "123"

	switch length := len(password); length >= 5 {
	case true:
		fmt.Println("Good Password")
	case false:
		fmt.Println("Password is too short")
	}
}
