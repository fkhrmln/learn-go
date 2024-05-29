package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM user;"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Id : %s\nName : %s\n", id, name)
	}
}
