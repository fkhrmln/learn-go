package main

import "fmt"

func main() {
	score := 50

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("Try Again")
	}

	password := "123"

	if length := len(password); length >= 5 {
		fmt.Println("Good Password")
	} else {
		fmt.Println("Password is too short")
	}
}
