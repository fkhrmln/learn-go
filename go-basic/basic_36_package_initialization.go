package main

import (
	"go-basic/example/database"
	_ "go-basic/example/socket"
)

func main() {
	database.GetDatabase()
}
