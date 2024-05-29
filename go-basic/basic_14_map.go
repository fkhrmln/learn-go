package main

import "fmt"

func main() {
	var name = map[string]string{
		"firstName":  "Fakhri",
		"middleName": "Maulana",
		"lastName":   "Ihsan",
	}

	fmt.Println(name["firstName"])
	fmt.Println(name["middleName"])
	fmt.Println(name["lastName"])
	fmt.Println(name)

	person := map[string]string{
		"name":    "Fakhri Maulana Ihsan",
		"age":     "20",
		"address": "Mutiara Gading Timur",
	}

	fmt.Println(person)
	fmt.Println(len(person))

	user := make(map[int]string)

	user[1] = "Fakhri"
	user[2] = "Rifky"
	user[3] = "Audry"
	user[4] = "Bobby"

	delete(user, 4)

	fmt.Println(user)
}
