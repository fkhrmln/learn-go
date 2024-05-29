package main

import "fmt"

func main() {
	const name string = "Fakhri Maulana Ihsan"

	// name = "Rifky Fediansyah" // Error

	fmt.Println(name)

	const (
		fistName   = "Fakhri"
		middleName = "Maulana"
		lastName   = "Ihsan"
	)
}
