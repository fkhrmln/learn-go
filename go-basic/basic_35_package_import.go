package main

import (
	"fmt"
	"go-basic/example/version"
)

func main() {
	// fmt.Println(helper.version) // Private

	fmt.Println(version.GetVersion())
}
