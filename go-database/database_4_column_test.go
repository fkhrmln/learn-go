package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestInsertColumn(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := fmt.Sprintf("INSERT INTO user (id, name, email, balance, rating, married, birthdate) VALUES ('%s', '%s', '%s', %d, %g, %t, '%s');", "00001", "Fakhri Maulana Ihsan", "fakhrimaulanaihsan@gmail.com", 1000000, 95.0, false, "2002-08-11")

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert User")
}

func TestSelectColumn(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, married, birthdate, created_at FROM user;"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id, name   string
			email      sql.NullString
			balance    int64
			rating     float64
			married    bool
			birthdate  sql.NullTime
			created_at time.Time
		)

		err := rows.Scan(&id, &name, &email, &balance, &rating, &married, &birthdate, &created_at)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)

		if email.Valid {
			fmt.Println("Email :", email.String)
		} else {
			fmt.Println("Email :", nil)
		}

		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		fmt.Println("Married :", married)

		if birthdate.Valid {
			fmt.Println("Birthdate :", birthdate)
		} else {
			fmt.Println("Birthdate :", nil)
		}

		fmt.Println("Created At :", created_at)
	}
}
