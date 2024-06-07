package entity

import "time"

type Student struct {
	ID        string    `gorm:"column:id;primaryKey;<-:create"`
	Name      string    `gorm:"column:name"`
	NPM       int       `gorm:"column:nim"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (student *Student) TableName() string {
	return "student"
}
