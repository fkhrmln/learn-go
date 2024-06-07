package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"testing"
)

func TestOneToMany(t *testing.T) {
	db := GetConnection()

	posts := []entity.Post{
		{
			ID:     "00001",
			UserID: "00001",
			Title:  "Title 1",
			Body:   "Body 1",
		},
		{
			ID:     "00002",
			UserID: "00001",
			Title:  "Title 2",
			Body:   "Body 2",
		},
		{
			ID:     "00003",
			UserID: "00001",
			Title:  "Title 3",
			Body:   "Body 3",
		},
	}

	err := db.Create(&posts).Error

	if err != nil {
		panic(err)
	}

	user := entity.User{}

	err = db.Model(&user).Preload("Post").Where("user.id = ?", "00001").Take(&user).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	posts = []entity.Post{}

	err = db.Model(&posts).Preload("User").Where("user_id = ?", "00001").Find(&posts).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(posts)

	posts = []entity.Post{}

	err = db.Model(&posts).Joins("User").Where("user_id = ?", "00001").Find(&posts).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(posts)
}
