package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func runOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			once.Do(runOnce)
		}()
	}

	group.Wait()

	fmt.Println(counter)
}
