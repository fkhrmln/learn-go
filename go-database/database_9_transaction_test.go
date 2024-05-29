package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment (username, comment) VALUES (?, ?);"

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	usernames := []string{"Fakhri Maulana Ihsan", "Rifky Ferdiansyah", "Audry Elgalia"}

	for _, username := range usernames {
		_, err := tx.ExecContext(ctx, query, username, "Hello World")

		if err != nil {
			panic(err)
		}

		fmt.Println(username, "Comment Successfully")
	}

	err = tx.Commit()

	if err != nil {
		panic(err)
	}
}
