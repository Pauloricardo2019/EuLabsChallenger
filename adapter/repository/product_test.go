package repository

import (
	"context"
	"eulabs_challenger/internal/config"
	"eulabs_challenger/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var logger *zap.Logger
var envPath = "../../dev.env"

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestProductRepository_Create(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	assert.NoError(t, err)

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productRepository := NewProductRepository(db, logger)

	product, err := productRepository.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, product.ID > 0)

}

func TestProductRepository_GetCount(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	assert.NoError(t, err)

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productRepository := NewProductRepository(db, logger)

	product, err := productRepository.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, product.ID > 0)

	count, err := productRepository.GetCount(ctx)
	assert.NoError(t, err)
	assert.True(t, count > 0)

}

func TestProductRepository_GetByID(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	assert.NoError(t, err)

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productRepository := NewProductRepository(db, logger)

	product, err := productRepository.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, product.ID > 0)

	found, productFound, err := productRepository.GetByID(ctx, product.ID)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.True(t, productFound.ID == productToCreate.ID)

}

func TestProductRepository_GetAll(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	assert.NoError(t, err)

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	limit := 10
	offset := 0

	productRepository := NewProductRepository(db, logger)

	product, err := productRepository.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, product.ID > 0)

	products, err := productRepository.GetAll(ctx, limit, offset)
	assert.NoError(t, err)
	assert.True(t, len(products) > 0)

}

func TestProductRepository_Update(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	assert.NoError(t, err)

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productRepository := NewProductRepository(db, logger)

	productCreated, err := productRepository.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, productToCreate.ID > 0)

	productCreated.Name = "test_product_change"
	productCreated.Description = "test_description_change"
	productCreated.Price = 89.99

	changeProduct, err := productRepository.Update(ctx, productCreated)
	assert.NoError(t, err)
	assert.Equal(t, "test_product_change", changeProduct.Name)
	assert.Equal(t, "test_description_change", changeProduct.Description)
	assert.Equal(t, 89.99, changeProduct.Price)

}

func TestProductRepository_Delete(t *testing.T) {
	ctx := context.Background()

	cfg := config.NewConfig(logger).GetConfig(envPath)
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	assert.NoError(t, err)

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productRepository := NewProductRepository(db, logger)

	product, err := productRepository.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, product.ID > 0)

	err = productRepository.Delete(ctx, product.ID)
	assert.NoError(t, err)

	found, _, err := productRepository.GetByID(ctx, product.ID)
	assert.NoError(t, err)
	assert.False(t, found)
}
