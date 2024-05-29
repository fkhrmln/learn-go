package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Fakhri", "Rifky", "Audry"}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Contains(names, "Fakhri"))
	fmt.Println(slices.Index(names, "Fakhri"))
	fmt.Println(slices.Max(numbers))
	fmt.Println(slices.Min(numbers))
}
