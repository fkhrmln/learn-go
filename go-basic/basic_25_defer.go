package main

import "fmt"

func start() {
	defer finish()

	fmt.Println("Start Application")
}

func finish() {
	fmt.Println("Finish Application")
}

func main() {
	start()
}
