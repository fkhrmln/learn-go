package main

import "fmt"

func main() {
	value32 := 50000
	value64 := int64(value32)
	value16 := int16(value32) // Number Overflow

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)

	name := "Fakhri Maulana Ihsan"
	char := name[0]

	fmt.Println(char)
	fmt.Println(string(char))
}
