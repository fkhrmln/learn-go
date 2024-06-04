package main

import (
	"fmt"
	"go-dependency-injection/helper"
	"go-dependency-injection/injector"
)

func main() {
	server := injector.InitializeServer()

	err := server.ListenAndServe()

	helper.PanicError(err)

	fmt.Println("Server is running on", server.Addr)
}
