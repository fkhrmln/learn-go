package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO user (id, name) VALUES ('00001', 'Fakhri Maulana Ihsan');"

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert User")
}
