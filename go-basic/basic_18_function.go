package main

import "fmt"

func sayHelloWorld() {
	fmt.Println("Hello World")
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sum(a int, b int) int {
	return a + b
}

func getNameAndAge(name string, age int) (string, int) {
	return name, age
}

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Fakhri"
	middleName = "Maulana"
	lastName = "Ihsan"

	return firstName, middleName, lastName
}

func main() {
	sayHelloWorld()

	sayHello("Fakhri")

	result := sum(5, 5)

	fmt.Println(result)

	nameOne, ageOne := getNameAndAge("Fakhri Maulana Ihsan", 20)

	println(nameOne, ageOne)

	nameTwo, _ := getNameAndAge("Fakhri Maulana Ihsan", 20)

	println(nameTwo)

	firstName, middleName, lastName := getFullName()

	println(firstName, middleName, lastName)
}
