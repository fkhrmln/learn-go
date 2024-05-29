package main

import "fmt"

type Account struct {
	Name string
}

func (account Account) ChangeName(name string) {
	account.Name = name
}

func (account *Account) ChangeNamePointer(name string) {
	account.Name = name
}

func main() {
	accountOne := Account{"Fakhri Maulana Ihsan"}

	accountOne.ChangeName("Rifky Ferdiansyah")

	fmt.Println(accountOne)

	accountTwo := Account{"Fakhri Maulana Ihsan"}

	accountTwo.ChangeNamePointer("Rifky Ferdiansyah")

	fmt.Println(accountTwo)
}
