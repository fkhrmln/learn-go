package entity

import "time"

type Post struct {
	ID        string    `gorm:"column:id;primaryKey;<-:create"`
	UserID    string    `gorm:"column:user_id"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (post *Post) TableName() string {
	return "post"
}
