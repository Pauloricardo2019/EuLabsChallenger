package repository

import (
	"context"
	"errors"
	"eulabs_challenger/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type productRepository struct {
	baseRepo *BaseRepository
	logger   *zap.Logger
}

func NewProductRepository(db *gorm.DB, logger *zap.Logger) *productRepository {
	baseRepo := NewBaseRepository(db)
	return &productRepository{
		baseRepo: baseRepo,
		logger:   logger,
	}
}

func (p *productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	p.logger.Info("Repository: Creating product")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(product).Error; err != nil {
		return nil, err
	}

	p.logger.Debug("Product", zap.Any("product", product))
	p.logger.Info("Repository: Product created")
	return product, nil
}

func (p *productRepository) GetCount(ctx context.Context) (int64, error) {
	p.logger.Info("Repository: Getting count")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return 0, err
	}

	products := make([]model.Product, 0)
	var count int64

	if err = conn.Find(&products).Count(&count).Error; err != nil {
		return 0, err
	}

	p.logger.Debug("Count", zap.Int64("count", count))
	p.logger.Info("Repository: Count gotten")

	return count, nil
}

func (p *productRepository) GetByID(ctx context.Context, id uint64) (bool, *model.Product, error) {
	p.logger.Info("Repository: Getting product by ID")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	product := &model.Product{}

	if err = conn.Where(&model.Product{
		ID: id,
	}).First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, product, nil
		}
		return false, nil, err
	}
	p.logger.Debug("Product", zap.Any("product", product), zap.Uint64("id", id), zap.Bool("found", true))
	p.logger.Info("Repository: Product gotten by ID")

	return true, product, nil
}

func (p *productRepository) GetAll(ctx context.Context, limit, offset int) ([]model.Product, error) {
	p.logger.Info("Repository: Getting all products")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]model.Product, 0)

	if err = conn.
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Products", zap.Any("products", products), zap.Int("limit", limit), zap.Int("offset", offset))
	p.logger.Info("Repository: Products gotten")

	return products, nil
}

func (p *productRepository) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	p.logger.Info("Repository: Updating product")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Debug().Save(product).Error; err != nil {
		return nil, err
	}
	p.logger.Debug("Product", zap.Any("product", product))
	p.logger.Info("Repository: Product updated")

	return product, nil
}

func (p *productRepository) Delete(ctx context.Context, id uint64) error {
	p.logger.Info("Repository: Deleting product")
	conn, err := p.baseRepo.getConnection(ctx)
	if err != nil {
		return err
	}

	p.logger.Debug("ID", zap.Uint64("id", id))
	if err = conn.Delete(&model.Product{
		ID: id,
	}).Error; err != nil {
		return err
	}

	p.logger.Info("Repository: Product deleted")

	return nil
}
