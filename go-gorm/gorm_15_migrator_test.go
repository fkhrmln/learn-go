package gogorm

import (
	"go-gorm/entity"
	"testing"
)

func TestMigrator(t *testing.T) {
	db := GetConnection()

	student := entity.Student{}

	err := db.Migrator().AutoMigrate(&student)

	if err != nil {
		panic(err)
	}
}
