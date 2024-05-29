package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestSQLParameter(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"

	password := "12345"

	query := "SELECT username FROM admin WHERE username = ? AND password = ? LIMIT 1;"

	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Login Successfully")
	} else {
		fmt.Println("Login Failed")
	}
}
