package main

import "fmt"

func runApplication(err bool) {
	defer endApplication()

	if err {
		panic("Ups Error")
	}

	fmt.Println("Run Application")
}

func endApplication() {
	fmt.Println("End Application")

	message := recover()

	fmt.Println(message)
}

func main() {
	runApplication(true)

	fmt.Println("Hello World")
}
