package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (personSlice PersonSlice) Len() int {
	return len(personSlice)
}

func (personSlice PersonSlice) Less(i int, j int) bool {
	return personSlice[i].Age < personSlice[j].Age
}

func (personSlice PersonSlice) Swap(i int, j int) {
	// temp := personSlice[i]
	// personSlice[i] = personSlice[j]
	// personSlice[j] = temp

	personSlice[i], personSlice[j] = personSlice[j], personSlice[i]
}

func main() {
	persons := []Person{
		{"Fakhri Maulana Ihsan", 20},
		{"Rifky Ferdiasnyah", 25},
		{"Audry Elgalia", 15},
	}

	fmt.Println(persons)

	sort.Sort(PersonSlice(persons))

	fmt.Println(persons)
}
