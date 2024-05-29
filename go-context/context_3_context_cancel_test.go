package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func counterCancel(ctx context.Context) chan int {
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

func TestContextCancel(t *testing.T) {
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	ctxParent := context.Background()

	ctx, cancel := context.WithCancel(ctxParent)

	channel := counterCancel(ctx)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	for data := range channel {
		fmt.Println(data)

		if data == 10 {
			break
		}
	}

	cancel()

	time.Sleep(1 * time.Second)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}
