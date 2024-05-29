package main

import "fmt"

func main() {
	var name string
	name = "Fakhri Maulana Ihsan"

	fmt.Println(name)

	// name = 1 // Error

	var age int = 20

	fmt.Println(age)

	address := "Mutiara Gading Timur"

	fmt.Println(address)

	var (
		firstName  = "Fakhri"
		middleName = "Maulana"
		lastName   = "Ihsan"
	)

	fmt.Println(firstName, middleName, lastName)
}
