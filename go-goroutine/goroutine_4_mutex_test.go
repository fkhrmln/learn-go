package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	sync.RWMutex
	Balance int
}

func (bankAccount *BankAccount) AddBalance(amount int) {
	bankAccount.RWMutex.Lock()

	bankAccount.Balance += amount

	bankAccount.RWMutex.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	bankAccount.RWMutex.RLock()

	balance := bankAccount.Balance

	bankAccount.RWMutex.RUnlock()

	return balance
}

func TestMutex(t *testing.T) {
	counter := 0

	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 1000; j++ {
				mutex.Lock()

				counter = counter + 1

				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(counter)
}

func TestRWMutext(t *testing.T) {
	account := BankAccount{
		Balance: 0,
	}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				account.AddBalance(1000)

				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final Balance", account.GetBalance())
}
