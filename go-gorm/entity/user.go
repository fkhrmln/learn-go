package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"column:id;primaryKey;<-:create"`
	Name        string         `gorm:"column:name"`
	Bank        Bank           `gorm:"foreignKey:user_id;refereces:id"`
	Post        []Post         `gorm:"foreignKey:user_id;refereces:id"`
	UserItem    []Item         `gorm:"many2many:user_item;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:item_id"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
	Information string         `gorm:"-"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	return nil
}

func (user *User) AfterCreate(db *gorm.DB) error {
	return nil
}

func (user *User) BeforeUpdate(db *gorm.DB) error {
	return nil
}

func (user *User) AfterUpdate(db *gorm.DB) error {
	return nil
}

func (user *User) BeforeDelete(db *gorm.DB) error {
	return nil
}

func (user *User) AfterDelete(db *gorm.DB) error {
	return nil
}

func (user *User) AfterFind(db *gorm.DB) error {
	return nil
}
