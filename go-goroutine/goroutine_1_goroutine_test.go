package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestGoroutine(t *testing.T) {
	go HelloWorld()

	fmt.Println("Finish")

	time.Sleep(1 * time.Second)
}

func TestLoopWithGoroutine(t *testing.T) {
	for i := 1; i <= 100000; i++ {
		go func() {
			fmt.Println("Hello World -", i)
		}()
	}

	time.Sleep(5 * time.Second)
}
