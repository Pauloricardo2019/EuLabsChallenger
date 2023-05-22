package service

import (
	"context"
	"eulabs_challenger/internal/model"
	"go.uber.org/zap"
)

type productService struct {
	productRepository productRepository
	logger            *zap.Logger
}

func NewProductService(productRepository productRepository, logger *zap.Logger) *productService {
	return &productService{
		productRepository: productRepository,
		logger:            logger,
	}
}

func (p *productService) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	p.logger.Info("Service: Creating product")
	p.logger.Debug("Product", zap.Any("product", product))
	return p.productRepository.Create(ctx, product)
}

func (p *productService) GetCount(ctx context.Context) (int64, error) {
	p.logger.Info("Service: Getting count")
	return p.productRepository.GetCount(ctx)
}

func (p *productService) GetByID(ctx context.Context, id uint64) (bool, *model.Product, error) {
	p.logger.Info("Service: Getting product by ID")
	p.logger.Debug("ID", zap.Uint64("id", id))
	return p.productRepository.GetByID(ctx, id)
}

func (p *productService) GetAll(ctx context.Context, limit, offset int) ([]model.Product, error) {
	p.logger.Info("Service: Getting all products")
	p.logger.Debug("Limit", zap.Int("limit", limit), zap.Int("offset", offset))
	return p.productRepository.GetAll(ctx, limit, offset)
}

func (p *productService) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	p.logger.Info("Service: Updating product")
	p.logger.Debug("Product", zap.Any("product", product))
	return p.productRepository.Update(ctx, product)
}

func (p *productService) Delete(ctx context.Context, id uint64) error {
	p.logger.Info("Service: Deleting product")
	p.logger.Debug("ID", zap.Uint64("id", id))
	return p.productRepository.Delete(ctx, id)
}
