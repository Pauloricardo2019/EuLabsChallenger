package service

import (
	"context"
	"eulabs_challenger/internal/model"
)

type productService struct {
	productRepository productRepository
}

func NewProductService(productRepository productRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p *productService) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	return p.productRepository.Create(ctx, product)
}

func (p *productService) GetByID(ctx context.Context, id uint64) (bool, *model.Product, error) {
	return p.productRepository.GetByID(ctx, id)
}

func (p *productService) GetAll(ctx context.Context, pagination *model.Pagination) ([]model.Product, error) {
	return p.productRepository.GetAll(ctx, pagination)
}

func (p *productService) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	return p.productRepository.Update(ctx, product)
}

func (p *productService) Delete(ctx context.Context, id uint64) error {
	return p.productRepository.Delete(ctx, id)
}
