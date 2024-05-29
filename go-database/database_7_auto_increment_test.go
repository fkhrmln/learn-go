package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment (username, message) VALUES (?, ?);"

	result, err := db.ExecContext(ctx, query, "Fakhri Maulana Ihsan", "Hello World")

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println(lastInsertId)
}
