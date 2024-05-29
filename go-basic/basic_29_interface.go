package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayName(hasName HasName) {
	fmt.Println("My name is", hasName.GetName())
}

type Admin struct {
	Name string
}

func (admin Admin) GetName() string {
	return admin.Name
}

type User struct {
	Name string
}

func (user User) GetName() string {
	return user.Name
}

func anything() any {
	// return "Hello"

	// return 1

	return true
}

func main() {
	admin := Admin{"Fakhri Maulana Ihsan"}

	SayName(admin)

	user := User{"Rifky Ferdiansyah"}

	SayName(user)

	fmt.Println(anything())
}
