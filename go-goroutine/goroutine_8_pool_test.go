package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "Empty"
		},
	}

	group := sync.WaitGroup{}

	pool.Put("Fakhri")
	pool.Put("Rifky")
	pool.Put("Audry")

	for i := 1; i <= 10; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			data := pool.Get()

			fmt.Println(data)

			pool.Put(data)
		}()
	}

	group.Wait()

	fmt.Println("Finish")
}
