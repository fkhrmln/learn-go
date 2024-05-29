package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Fakhri"
	names[1] = "Rifky"
	names[2] = "Audry"

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[0], names[1], names[2])

	var numbers = [...]int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
}
