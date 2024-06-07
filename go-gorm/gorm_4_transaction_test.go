package gogorm

import (
	"go-gorm/entity"
	"testing"

	"gorm.io/gorm"
)

func TestTransactionCommit(t *testing.T) {
	db := GetConnection()

	err := db.Transaction(func(tx *gorm.DB) error {
		userOne := entity.User{
			ID:   "00001",
			Name: "Fakhri Maulana Ihsan",
		}

		err := tx.Create(&userOne).Error

		if err != nil {
			return err
		}

		userTwo := entity.User{
			ID:   "00002",
			Name: "Rifky Ferdiansyah",
		}

		err = tx.Create(&userTwo).Error

		if err != nil {
			return err
		}

		userThree := entity.User{
			ID:   "00003",
			Name: "Audry Elgalia",
		}

		err = tx.Create(&userThree).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func TestTransactionRollback(t *testing.T) {
	db := GetConnection()

	err := db.Transaction(func(tx *gorm.DB) error {
		userFour := entity.User{
			ID:   "00004",
			Name: "Bobby Pratama",
		}

		err := tx.Create(&userFour).Error

		if err != nil {
			return err
		}

		userFive := entity.User{
			ID:   "00001",
			Name: "Fakhri Maulana Ihsan",
		}

		err = tx.Create(&userFive).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
