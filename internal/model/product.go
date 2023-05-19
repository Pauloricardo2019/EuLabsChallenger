package model

import "time"

type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment;column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Price       float64   `gorm:"column:price"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli;column:updated_at"`
}

func (Product) TableName() string {
	return "products"
}
