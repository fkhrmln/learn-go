package gogorm

import (
	"go-gorm/entity"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestLock(t *testing.T) {
	db := GetConnection()

	err := db.Transaction(func(tx *gorm.DB) error {
		product := entity.Product{}

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", 1).Take(&product).Error

		if err != nil {
			return err
		}

		product.Price = 1000000

		err = tx.Save(&product).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
