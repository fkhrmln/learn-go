package main

import "fmt"

type Customer struct {
	Name    string
	Age     int
	Address string
}

func main() {
	customerOne := Customer{"Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	customerTwo := customerOne

	customerTwo.Name = "Rifky Ferdiansyah"

	fmt.Println(customerOne)
	fmt.Println(customerTwo)

	customerThree := Customer{"Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	customerFour := &customerThree

	customerFour.Name = "Rifky Ferdiasnyah"

	fmt.Println(customerThree)
	fmt.Println(customerFour)

	customerFive := Customer{"Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	customerSix := &customerFive

	customerSix.Name = "Rifky Ferdiansyah"

	fmt.Println(customerFive)
	fmt.Println(customerSix)

	customerSix = &Customer{"Audry Elgalia", 20, "Rawa Kalong"}

	fmt.Println(customerFive)
	fmt.Println(customerSix)

	customerSeven := Customer{"Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	customerEight := &customerSeven

	customerEight.Name = "Rifky Ferdiansyah"

	fmt.Println(customerSeven)
	fmt.Println(customerEight)

	*customerEight = Customer{"Audry Elgalia", 20, "Rawa Kalong"}

	fmt.Println(customerSeven)
	fmt.Println(customerEight)

	customerNine := new(Customer)

	customerTen := customerNine

	customerTen.Name = "Fakhri Maulana Ihsan"

	fmt.Println(customerNine)
	fmt.Println(customerTen)
}
