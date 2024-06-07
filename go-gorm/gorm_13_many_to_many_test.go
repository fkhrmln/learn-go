package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"testing"
)

func TestManyToMany(t *testing.T) {
	db := GetConnection()

	item := entity.Item{
		ID:    "00001",
		Name:  "Item 1",
		Price: 1000000,
	}

	err := db.Create(&item).Error

	if err != nil {
		panic(err)
	}

	err = db.Table("user_item").Create(map[string]interface{}{
		"user_id": "00001",
		"item_id": "00001",
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Table("user_item").Create(map[string]interface{}{
		"user_id": "00002",
		"item_id": "00001",
	}).Error

	if err != nil {
		panic(err)
	}

	err = db.Table("user_item").Create(map[string]interface{}{
		"user_id": "00003",
		"item_id": "00001",
	}).Error

	if err != nil {
		panic(err)
	}

	user := entity.User{}

	err = db.Preload("UserItem").Where("id = ?", "00001").Take(&user).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	item = entity.Item{}

	err = db.Preload("ItemUser").Where("id = ?", "00001").Take(&item).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(item)
}
