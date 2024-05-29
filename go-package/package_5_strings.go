package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Fakhri Maulana Ihsan"))
	fmt.Println(strings.ToLower("Fakhri Maulana Ihsan"))
	fmt.Println(strings.Contains("Fakhri Maulana Ihsan", "Maulana"))
	fmt.Println(strings.Split("Fakhri Maulana Ihsan", " "))
	fmt.Println(strings.Trim("   Fakhri Maulana Ihsan   ", " "))
	fmt.Println(strings.ReplaceAll("Fakhri Maulana Ihsan", "Maulana", "Muhammad"))
}
