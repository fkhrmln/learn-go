package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "127.0.0.1", "Host For Database")
	port := flag.String("port", "5432", "Port For Database")
	username := flag.String("username", "root", "Username For Database")
	password := flag.String("password", "root", "Password For Database")

	flag.Parse()

	fmt.Println(*host, *port, *username, *password)
}
