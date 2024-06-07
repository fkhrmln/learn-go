package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"strconv"
	"testing"
)

func TestCreate(t *testing.T) {
	db := GetConnection()

	user := entity.User{
		ID:   "00001",
		Name: "Fakhri Maulana Ihsan",
	}

	rowsAffected := db.Create(&user).RowsAffected

	fmt.Println(rowsAffected)
}

func TestCreateBatch(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	for i := 1; i <= 100; i++ {
		user := entity.User{
			ID:   strconv.Itoa(i),
			Name: "User - " + strconv.Itoa(i),
		}

		users = append(users, user)
	}

	rowsAffected := db.Create(&users).RowsAffected

	fmt.Println(rowsAffected)
}

func TestCreateInBatch(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	for i := 1; i <= 100; i++ {
		user := entity.User{
			ID:   strconv.Itoa(i),
			Name: "User - " + strconv.Itoa(i),
		}

		users = append(users, user)
	}

	rowsAffected := db.CreateInBatches(&users, 10).RowsAffected

	fmt.Println(rowsAffected)
}
