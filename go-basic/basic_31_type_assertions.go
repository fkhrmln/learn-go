package main

import "fmt"

func getStatus() interface{} {
	return "Success"
}

func getStatusCode() interface{} {
	return 404
}

func main() {
	status := getStatus()

	statusString := status.(string)

	fmt.Println(statusString)

	// statusInt := status.(int) // Panic

	statusCode := getStatusCode()

	switch statusCodeType := statusCode.(type) {
	case string:
		fmt.Println("String :", statusCodeType)
	case int:
		fmt.Println("Int :", statusCodeType)
	default:
		fmt.Println("Unknown Type")
	}
}
