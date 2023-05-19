package v1

import (
	"context"
	"eulabs_challenger/internal/dto"
)

type productFacade interface {
	CreateProduct(ctx context.Context, product *dto.CreateProductRequest) (*dto.CreateProductResponse, error)
	GetByIDProduct(ctx context.Context, id uint64) (*dto.GetByProductIDResponse, error)
	GetAllProducts(ctx context.Context, limit, offset int) (*dto.GetAllProductsResponse, error)
	UpdateProduct(ctx context.Context, productID uint64, product *dto.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id uint64) error
}
