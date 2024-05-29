package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (userBalance *UserBalance) ChangeBalance(amount int) {
	userBalance.Balance += amount
}

func TransferBalance(from *UserBalance, to *UserBalance, amount int) {
	from.Mutex.Lock()

	fmt.Println("Lock", from.Name)

	from.ChangeBalance(-amount)

	time.Sleep(1 * time.Second)

	to.Mutex.Lock()

	fmt.Println("Lock", to.Name)

	from.ChangeBalance(amount)

	time.Sleep(1 * time.Second)

	from.Mutex.Unlock()
	to.Mutex.Unlock()
}

func TestDeadlock(t *testing.T) {
	userOne := UserBalance{
		Name:    "Fakhri Maulana",
		Balance: 1000000,
	}

	userTwo := UserBalance{
		Name:    "Rifky Ferdiansyah",
		Balance: 1000000,
	}

	go TransferBalance(&userOne, &userTwo, 100000)

	go TransferBalance(&userTwo, &userOne, 100000)

	time.Sleep(5 * time.Second)

	fmt.Printf("%s Balance : %d\n", userOne.Name, userOne.Balance)
	fmt.Printf("%s Balance : %d\n", userTwo.Name, userTwo.Balance)
}
