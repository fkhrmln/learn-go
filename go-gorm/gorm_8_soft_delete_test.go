package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"testing"
)

func TestSoftDelete(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	err := db.Where("id = ?", "00003").Delete(&user).Error

	if err != nil {
		panic(err)
	}

	users := []entity.User{}

	err = db.Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	err = db.Unscoped().Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
