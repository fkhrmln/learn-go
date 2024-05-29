package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func (person Person) SayHello(name string) {
	fmt.Printf("Hello %s My name is %s\n", name, person.Name)
}

func main() {
	var personOne Person

	personOne.Name = "Fakhri Maulana Ihsan"
	personOne.Age = 20
	personOne.Address = "Mutiara Gading Timur"

	fmt.Println(personOne)
	fmt.Println(personOne.Name)
	fmt.Println(personOne.Age)
	fmt.Println(personOne.Address)

	personTwo := Person{
		Name:    "Rifky Ferdiansyah",
		Age:     20,
		Address: "Bekasi Timur Regency",
	}

	fmt.Println(personTwo)
	fmt.Println(personTwo.Name)
	fmt.Println(personTwo.Age)
	fmt.Println(personTwo.Address)

	personThree := Person{"Audry Elgalia", 20, "Rawa Kalong"}

	fmt.Println(personThree)
	fmt.Println(personThree.Name)
	fmt.Println(personThree.Age)
	fmt.Println(personThree.Address)

	personOne.SayHello("Audry Elgalia")
}
