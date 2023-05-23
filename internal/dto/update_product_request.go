package dto

import "eulabs_challenger/internal/model"

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
} // @name UpdateProductRequest

func (u *UpdateProductRequest) ConvertToProductVO() *model.Product {
	return &model.Product{
		Name:        u.Name,
		Description: u.Description,
		Price:       u.Price,
	}
}
