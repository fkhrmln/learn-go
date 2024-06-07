package gogorm

import (
	"go-gorm/entity"
	"testing"
)

func TestGORMModel(t *testing.T) {
	db := GetConnection()

	product := entity.Product{
		Name:  "Product 1",
		Price: 1000000,
	}

	err := db.Create(&product).Error

	if err != nil {
		panic(err)
	}
}
