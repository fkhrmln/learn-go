package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(`Fakhri \w+ Ihsan`)

	fmt.Println(pattern.MatchString("Fakhri Maulana Ihsan"))
	fmt.Println(pattern.MatchString("Fakhri Darwanto Ihsan"))
	fmt.Println(pattern.MatchString("Fakhri Hudriati Ihsan"))
}
