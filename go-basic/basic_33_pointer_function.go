package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Address string
}

func changeAddressToBTR(student Student) {
	student.Address = "Bekasi Timur Regency"
}

func changeAddressToBTRPointer(student *Student) {
	student.Address = "Bekasi Timur Regency"
}

func main() {
	studentOne := Student{"Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	changeAddressToBTR(studentOne)

	fmt.Println(studentOne)

	studentTwo := Student{"Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	changeAddressToBTRPointer(&studentTwo)

	fmt.Println(studentTwo)
}
