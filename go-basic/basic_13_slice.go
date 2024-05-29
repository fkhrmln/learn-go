package main

import "fmt"

func main() {
	months := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println(months[7:8])

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	numberSlice := numbers[2:7]

	fmt.Println(numberSlice)
	fmt.Println(len(numberSlice))
	fmt.Println(cap(numberSlice))

	names := [...]string{"Fakhri", "Maulana", "Ihsan"}

	nameSliceOne := names[1:3]

	fmt.Println(nameSliceOne)

	nameSliceOne[0] = "Rifky"
	nameSliceOne[1] = "Audry"

	fmt.Println(nameSliceOne)

	fmt.Println(names)

	nameSliceTwo := append(nameSliceOne, "Bobby")

	fmt.Println(nameSliceTwo)

	fmt.Println(nameSliceOne)

	fmt.Println(names)

	sliceOne := make([]string, 3, 5)

	sliceOne[0] = "Fakhri"
	sliceOne[1] = "Rifky"
	sliceOne[2] = "Audry"

	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	sliceTwo := append(sliceOne, "Bobby")

	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo))
	fmt.Println(cap(sliceTwo))

	sliceThree := make([]string, len(sliceTwo), cap(sliceTwo))

	copy(sliceThree, sliceTwo)

	fmt.Println(sliceThree)
	fmt.Println(len(sliceThree))
	fmt.Println(cap(sliceThree))
}
