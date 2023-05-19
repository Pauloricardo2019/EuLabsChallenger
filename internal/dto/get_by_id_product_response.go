package dto

import (
	"eulabs_challenger/internal/model"
	"time"
)

type GetByProductIDResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func (g *GetByProductIDResponse) ParseFromProductVO(product *model.Product) {
	g.ID = product.ID
	g.Name = product.Name
	g.Description = product.Description
	g.Price = product.Price
	g.CreateAt = product.CreatedAt
	g.UpdateAt = product.UpdatedAt
}
