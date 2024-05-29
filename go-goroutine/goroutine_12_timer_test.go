package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	time := <-timer.C

	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	timer := time.After(5 * time.Second)

	time := <-timer

	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		defer group.Done()

		fmt.Println("Finish After 5 Second")
	})

	group.Wait()
}
