package dto

import "eulabs_challenger/internal/model"

type UpdateProductRequest struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (u *UpdateProductRequest) ConvertToProductVO() *model.Product {
	return &model.Product{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
		Price:       u.Price,
	}
}
