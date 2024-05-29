package main

import "fmt"

func createPersonMap(name string, age int, address string) map[string]any {
	if name == "" && age == 0 && address == "" {
		return nil
	}

	return map[string]any{
		"name":    name,
		"age":     age,
		"address": address,
	}
}

func main() {
	personOne := createPersonMap("Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur")

	if personOne == nil {
		fmt.Println("Empty")
	} else {
		fmt.Println(personOne)
	}

}
