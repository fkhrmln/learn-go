package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestSQLInjection(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"

	password := "12345"

	query := fmt.Sprintf("SELECT username FROM admin WHERE username = '%s' AND password = '%s' LIMIT 1;", username, password)

	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)

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
