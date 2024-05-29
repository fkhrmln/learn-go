package main

import "fmt"

type AgeRestrictions func(int) bool

func checkAgeRestrictions(age int, ageRestrictions AgeRestrictions) {
	if ageRestrictions(age) {
		fmt.Println("You can continue")

		return
	}

	fmt.Println("You can't continue")
}

func main() {
	checkAgeRestrictions(20, func(age int) bool {
		if age < 20 {
			return false
		}

		return true
	})
}
