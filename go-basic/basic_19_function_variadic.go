package main

import "fmt"

func sumSlice(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func sumVariadic(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	sumSliceResult := sumSlice([]int{1, 2, 3, 4, 5})

	fmt.Println(sumSliceResult)

	sumVariadicResult := sumVariadic(1, 2, 3, 4, 5)

	fmt.Println(sumVariadicResult)

	numbers := []int{1, 2, 3, 4, 5}

	sumVariadicSliceResult := sumVariadic(numbers...)

	fmt.Println(sumVariadicSliceResult)
}
