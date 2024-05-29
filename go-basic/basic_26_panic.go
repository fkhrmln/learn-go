package main

import "fmt"

func run(err bool) {
	defer end()

	if err {
		panic("Ups Error")
	}

	fmt.Println("Run Application")
}

func end() {
	fmt.Println("End Application")
}

func main() {
	run(true)
}
