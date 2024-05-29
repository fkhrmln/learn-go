package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(target *sync.Map, group *sync.WaitGroup, value int) {
	defer group.Done()

	target.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}

	group := &sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		group.Add(1)

		go addToMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Printf("%d : %d\n", key, value)

		return true
	})
}
