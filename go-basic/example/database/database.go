package database

import "fmt"

var database string

func init() {
	database = "PostgreSQL"
}

func GetDatabase() {
	fmt.Println(database, "Connected")
}
