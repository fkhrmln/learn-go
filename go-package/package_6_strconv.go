package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringToInt, err := strconv.Atoi("100")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stringToInt)
	}

	intToString := strconv.Itoa(100)

	fmt.Println(intToString)

	stringToBool, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stringToBool)
	}

	boolToString := strconv.FormatBool(true)

	fmt.Println(boolToString)
}
