package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	names := []string{"Fakhri", "Rifky", "Audry"}

	for i := 0; i < len(names); i++ {
		fmt.Printf("%d. %s\n", i, names[i])

	}

	for index, value := range names {
		fmt.Printf("%d. %s\n", index, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	person := map[string]string{
		"name":    "Fakhri Maulana Ihsan",
		"age":     "20",
		"address": "Mutiara Gading Timur",
	}

	for key, value := range person {
		fmt.Printf("%s : %s\n", key, value)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)

		if i == 5 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
