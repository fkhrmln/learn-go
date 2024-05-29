package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println(currentTime.Local())

	birthTime := time.Date(2002, time.August, 11, 0, 0, 0, 0, time.Local)

	fmt.Println(birthTime)

	parsedBirthTime, err := time.Parse(time.RFC3339, "2002-08-11T00:00:00Z")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(parsedBirthTime.Format("02 January 2006"))
	}

	fmt.Println(parsedBirthTime.Day())
	fmt.Println(parsedBirthTime.Month())
	fmt.Println(parsedBirthTime.Year())

	oneHour := 1 * time.Hour
	oneMinute := 1 * time.Minute
	oneSecond := 1 * time.Second

	fmt.Println(oneHour)
	fmt.Println(oneMinute)
	fmt.Println(oneSecond)
}
