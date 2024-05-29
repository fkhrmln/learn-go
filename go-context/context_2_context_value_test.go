package gocontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContextValue(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "B", "B")
	ctxC := context.WithValue(ctxA, "C", "C")

	ctxD := context.WithValue(ctxB, "D", "D")
	ctxE := context.WithValue(ctxB, "E", "E")

	ctxF := context.WithValue(ctxC, "F", "F")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)

	fmt.Println(ctxF.Value("A"))
	fmt.Println(ctxF.Value("C"))
	fmt.Println(ctxF.Value("F"))
}
