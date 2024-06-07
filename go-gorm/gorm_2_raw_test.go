package gogorm

import (
	"fmt"
	"testing"
)

func TestRawInsert(t *testing.T) {
	db := GetConnection()

	sql := "INSERT INTO user (id, name) VALUES (?, ?);"

	err := db.Exec(sql, "00001", "Fakhri Maulana Ihsan").Error

	if err != nil {
		panic(err)
	}
}

func TestRawSelect(t *testing.T) {
	type User struct {
		Id   string
		Name string
	}

	db := GetConnection()

	users := []User{}

	sql := "SELECT id, name FROM user;"

	err := db.Raw(sql).Scan(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
