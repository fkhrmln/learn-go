package main

import "fmt"

func main() {
	type Id string

	var myIdString string = "12345"

	var myId Id = Id(myIdString)

	fmt.Println(myIdString)
	fmt.Println(myId)
}
