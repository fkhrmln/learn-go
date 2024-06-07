package entity

import "time"

type Item struct {
	ID        string    `gorm:"column:id;primaryKey;<-:create"`
	Name      string    `gorm:"column:name"`
	Price     int       `gorm:"column:price"`
	ItemUser  []User    `gorm:"many2many:user_item;foreignKey:id;joinForeignKey:item_id;references:id;joinReferences:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (item *Item) TableName() string {
	return "item"
}
