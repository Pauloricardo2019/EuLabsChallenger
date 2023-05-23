package dto

import "eulabs_challenger/internal/model"

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
} // @name CreateProductRequest

func (c *CreateProductRequest) ConvertToProductVO() *model.Product {
	return &model.Product{
		Name:        c.Name,
		Description: c.Description,
		Price:       c.Price,
	}
}
