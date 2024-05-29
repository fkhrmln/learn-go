package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 1000; j++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(counter)
}
