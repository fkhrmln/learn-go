package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := 10
	d := 10
	e := 10

	result := a + b - c*d/e

	fmt.Println(result)

	result *= 10

	fmt.Println(result)

	result++
	result--

	fmt.Println(result)
}
