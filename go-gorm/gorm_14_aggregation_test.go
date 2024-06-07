package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"testing"
)

func TestAggregation(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	var count int64

	err := db.Model(&user).Count(&count).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(count)

	bank := entity.Bank{}

	type BankAggregation struct {
		SUM int64
		AVG float64
		MAX int64
		MIN int64
	}

	bankAggregation := BankAggregation{}

	err = db.Model(&bank).Select("SUM(balance) AS SUM", "AVG(balance) AS AVG", "MAX(balance) AS MAX", "MIN(balance) AS MIN").Joins("User").Group("User.id").Having("SUM(balance) >= 500000").Find(&bankAggregation).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(bankAggregation)
}
