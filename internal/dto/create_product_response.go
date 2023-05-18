package dto

import "eulabs_challenger/internal/model"

type CreateProductResponse struct {
	ID uint64 `json:"id"`
}

func (c *CreateProductResponse) ParseFromProductVO(product *model.Product) {
	c.ID = product.ID
}
