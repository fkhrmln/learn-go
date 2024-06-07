package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"testing"
)

func TestOneToOne(t *testing.T) {
	db := GetConnection()

	bank := entity.Bank{
		ID:      "00001",
		UserID:  "00001",
		Balance: 1000000,
	}

	err := db.Create(&bank).Error

	if err != nil {
		panic(err)
	}

	user := entity.User{}

	err = db.Model(&user).Joins("Bank").Where("user.id = ?", "00001").Take(&user).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	bank = entity.Bank{}

	err = db.Model(&bank).Joins("User").Where("user_id = ?", "00001").Take(&bank).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(bank)
}
