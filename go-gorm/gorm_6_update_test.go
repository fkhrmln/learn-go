package gogorm

import (
	"go-gorm/entity"
	"testing"
)

func TestUpdate(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	err := db.Where("id = ?", "00001").Take(&user).Error

	if err != nil {
		panic(err)
	}

	err = db.Model(&user).Update("name", "Bobby Pratama").Error

	if err != nil {
		panic(err)
	}

	err = db.Model(&user).Updates(map[string]interface{}{
		"name": "Fakhri Maulana Ihsan",
	}).Error

	if err != nil {
		panic(err)
	}
}
