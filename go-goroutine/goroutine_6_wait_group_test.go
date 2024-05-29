package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter int

func runGoroutine(value int, group *sync.WaitGroup) {
	defer group.Done()

	fmt.Println("Hello World", value)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)

		go runGoroutine(i, group)
	}

	group.Wait()

	fmt.Println("Finish")
}
