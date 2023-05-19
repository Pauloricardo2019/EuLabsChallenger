package dto

import (
	"eulabs_challenger/internal/model"
	"time"
)

type ProductPagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

type GetAllProductsResponse struct {
	Products   []Product         `json:"products"`
	Pagination ProductPagination `json:"pagination"`
}

type Product struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (g *GetAllProductsResponse) ParseFromProductVO(products []model.Product, limit, offset int, total int64) {
	g.Pagination.Limit = limit
	g.Pagination.Offset = offset
	g.Pagination.Total = total

	for _, product := range products {
		g.Products = append(g.Products, Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
}
