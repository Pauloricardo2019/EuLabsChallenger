package model

import "time"

type Product struct {
	ID          uint64  `gorm:"primary_key;auto_increment;column:id"`
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	Quantity    int     `gorm:"column:quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Product) TableName() string {
	return "products"
}
