package gogoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGOMAXPROCS(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			time.Sleep(3 * time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()

	fmt.Println("Total CPU :", totalCPU)

	runtime.GOMAXPROCS(10)

	totalThread := runtime.GOMAXPROCS(-1)

	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()

	fmt.Println("Total Goroutine :", totalGoroutine)

	group.Wait()
}
