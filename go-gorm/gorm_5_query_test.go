package gogorm

import (
	"fmt"
	"go-gorm/entity"
	"testing"
)

func TestQueryFirst(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	err := db.First(&user).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func TestQueryLast(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	err := db.Last(&user).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func TestQueryTake(t *testing.T) {
	db := GetConnection()

	user := entity.User{}

	err := db.Take(&user).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func TestQueryFind(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQuerySelect(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Select("name").Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryCondition(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Find(&users, "id = ?", "00001").Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryInlineCondition(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Where("id = ?", "00001").Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryAnd(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Where("id = ?", "00001").Where("name = ?", "Fakhri Mualana Ihsan").Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryOr(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Where("id = ?", "00001").Or("name = ?", "Fakhri Mualana Ihsan").Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryNot(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Not("id = ?", "00001").Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryStructCondition(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	type UsersCondition struct {
		Id string
	}

	usersCondition := UsersCondition{
		Id: "00001",
	}

	err := db.Where(usersCondition).Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryMapCondition(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	usersCondition := map[string]interface{}{
		"id": "00001",
	}

	err := db.Where(usersCondition).Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryOrderLimitOffset(t *testing.T) {
	db := GetConnection()

	users := []entity.User{}

	err := db.Order("id ASC").Limit(3).Offset(0).Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func TestQueryNonModel(t *testing.T) {
	db := GetConnection()

	type UserResponse struct {
		Id   string
		Name string
	}

	users := []UserResponse{}

	err := db.Model(entity.User{}).Select("id", "name").Find(&users).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
