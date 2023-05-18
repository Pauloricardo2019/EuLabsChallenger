package facade

import (
	"context"
	"eulabs_challenger/internal/model"
)

type productService interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.Product, error)
	GetAll(ctx context.Context, pagination *model.Pagination) ([]model.Product, error)
	Update(ctx context.Context, product *model.Product) (*model.Product, error)
	Delete(ctx context.Context, id uint64) error
}
