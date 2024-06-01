package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id      string `required:"true"`
	Name    string `required:"true" max:"50"`
	Age     int    `required:"true" max:"3"`
	Address string
}

func getFields(value interface{}) {
	valueType := reflect.TypeOf(value)

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)

		fmt.Printf("%s : %s : %s\n", field.Name, field.Type, field.Tag)
	}
}

func isValid(value interface{}) (valid bool) {
	valueType := reflect.TypeOf(value)

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)

		if field.Tag.Get("required") == "true" {
			valid = reflect.ValueOf(value).Field(i).Interface() != ""

			if valid == false {
				return valid
			}
		}
	}

	return valid
}

func main() {
	userOne := User{"", "Fakhri Maulana Ihsan", 20, "Mutiara Gading Timur"}

	getFields(userOne)

	fmt.Println(isValid(userOne))
}
