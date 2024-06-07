package gogorm

import (
	"go-gorm/entity"
	"testing"
)

func TestDelete(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	err := db.Where("id = ?", "00003").Delete(&user).Error

	if err != nil {
		panic(err)
	}
}
