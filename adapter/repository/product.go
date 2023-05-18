package repository

import (
	"context"
	"errors"
	"eulabs_challenger/internal/model"
	"gorm.io/gorm"
)

type productRepository struct {
	*BaseRepository
}

func NewProductRepository(db *gorm.DB) *productRepository {
	baseRepo := NewBaseRepository(db)
	return &productRepository{
		baseRepo,
	}
}

func (p *productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	conn, err := p.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productRepository) GetCount(ctx context.Context) (int64, error) {
	conn, err := p.getConnection(ctx)
	if err != nil {
		return 0, err
	}

	products := make([]model.Product, 0)
	var count int64

	if err = conn.Find(&products).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (p *productRepository) GetByID(ctx context.Context, id uint64) (bool, *model.Product, error) {
	conn, err := p.getConnection(ctx)
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

	return true, product, nil
}

func (p *productRepository) GetAll(ctx context.Context, limit, offset int) ([]model.Product, error) {
	conn, err := p.getConnection(ctx)
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

	return products, nil
}

func (p *productRepository) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	conn, err := p.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = conn.Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productRepository) Delete(ctx context.Context, id uint64) error {
	conn, err := p.getConnection(ctx)
	if err != nil {
		return err
	}

	return conn.Delete(&model.Product{
		ID: id,
	}).Error
}
