package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func counterTimeout(ctx context.Context) chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				channel <- counter

				time.Sleep(1 * time.Second)

				counter++
			}
		}
	}()

	return channel
}

func TestContextTimeout(t *testing.T) {
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	ctxParent := context.Background()

	ctx, cancel := context.WithTimeout(ctxParent, 5*time.Second)

	defer cancel()

	channel := counterTimeout(ctx)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}
