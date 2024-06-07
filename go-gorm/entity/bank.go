package entity

import "time"

type Bank struct {
	ID        string    `gorm:"column:id;primaryKey;<-:create"`
	UserID    string    `gorm:"column:user_id"`
	Balance   int       `gorm:"column:balance"`
	User      *User     `gorm:"foreignKey:user_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (bank *Bank) TableName() string {
	return "bank"
}
