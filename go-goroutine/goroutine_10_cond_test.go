package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var mutex = &sync.Mutex{}

var cond = sync.NewCond(mutex)

var group = sync.WaitGroup{}

func waitCond(value int) {
	defer group.Done()

	cond.L.Lock()

	cond.Wait()

	fmt.Println("Hello World -", value)

	cond.L.Unlock()
}

func TestWaitCond(t *testing.T) {
	for i := 1; i <= 10; i++ {
		group.Add(1)

		go waitCond(i)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)

			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)

	// 	cond.Broadcast()
	// }()

	group.Wait()
}
