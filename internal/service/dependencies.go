package service

import (
	"context"
	"eulabs_challenger/internal/model"
)

type productRepository interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
	GetCount(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.Product, error)
	GetAll(ctx context.Context, limit, offset int) ([]model.Product, error)
	Update(ctx context.Context, product *model.Product) (*model.Product, error)
	Delete(ctx context.Context, id uint64) error
}
