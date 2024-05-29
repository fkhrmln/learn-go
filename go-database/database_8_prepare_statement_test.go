package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment (username, message) VALUES (?, ?);"

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	usernames := []string{"Fakhri Maulana Ihsan", "Rifky Ferdiansyah", "Audry Elgalia"}

	for _, username := range usernames {
		_, err := stmt.ExecContext(ctx, username, "Hello World")

		if err != nil {
			panic(err)
		}

		fmt.Println(username, "Comment Successfully")
	}
}
