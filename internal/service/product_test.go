package service

import (
	"context"
	"eulabs_challenger/internal/mocks"
	"eulabs_challenger/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
	"time"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func TestProductService_Create(t *testing.T) {
	ctx := context.Background()

	productRepositoryMock := &mocks.ProductRepositoryMock{}

	productToCreate := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productCreated := &model.Product{
		ID:          1,
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productRepositoryMock.On("Create", ctx, mock.Anything).
		Return(
			productCreated,
			nil,
		)

	productService := NewProductService(productRepositoryMock, logger)

	productCreated, err := productService.Create(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, productCreated.ID == 1)

}

func TestProductService_GetCount(t *testing.T) {
	ctx := context.Background()

	productRepositoryMock := &mocks.ProductRepositoryMock{}

	productRepositoryMock.On("GetCount", ctx).
		Return(
			int64(1),
			nil,
		)

	productService := NewProductService(productRepositoryMock, logger)

	count, err := productService.GetCount(ctx)
	assert.NoError(t, err)
	assert.True(t, count == 1)

}

func TestProductService_GetByID(t *testing.T) {
	ctx := context.Background()

	productRepositoryMock := &mocks.ProductRepositoryMock{}

	idMock := uint64(1)

	productFound := &model.Product{
		ID:          1,
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productRepositoryMock.On("GetByID", ctx, idMock).
		Return(
			true,
			productFound,
			nil,
		)

	productService := NewProductService(productRepositoryMock, logger)

	found, productFound, err := productService.GetByID(ctx, idMock)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.True(t, productFound.ID == 1)

}

func TestProductService_GetAll(t *testing.T) {
	ctx := context.Background()

	productRepositoryMock := &mocks.ProductRepositoryMock{}

	productsFound := []model.Product{
		{
			ID:          1,
			Name:        "test_product",
			Description: "test_description",
			Price:       10.0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "test_product2",
			Description: "test_description2",
			Price:       99.90,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	productRepositoryMock.On("GetAll", ctx, mock.Anything, mock.Anything).
		Return(
			productsFound,
			nil,
		)

	productService := NewProductService(productRepositoryMock, logger)

	products, err := productService.GetAll(ctx, 1, 10)
	assert.NoError(t, err)
	assert.True(t, len(products) == 2)

}

func TestProductService_Update(t *testing.T) {
	ctx := context.Background()

	productRepositoryMock := &mocks.ProductRepositoryMock{}

	productToUpdate := &model.Product{
		ID:          1,
		Name:        "test_product_change",
		Description: "test_description_change",
		Price:       10.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productUpdated := &model.Product{
		ID:          1,
		Name:        "test_product_change",
		Description: "test_description_change",
		Price:       10.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productRepositoryMock.On("Update", ctx, mock.Anything).
		Return(
			productUpdated,
			nil,
		)

	productService := NewProductService(productRepositoryMock, logger)

	productUpdated, err := productService.Update(ctx, productToUpdate)
	assert.NoError(t, err)
	assert.True(t, productUpdated.ID == 1)
	assert.True(t, productUpdated.Name == "test_product_change")
	assert.True(t, productUpdated.Description == "test_description_change")
	assert.True(t, productUpdated.Price == 10.0)
}

func TestProductService_Delete(t *testing.T) {
	ctx := context.Background()

	productRepositoryMock := &mocks.ProductRepositoryMock{}

	productID := uint64(1)

	productRepositoryMock.On("Delete", ctx, productID).
		Return(
			nil,
		)

	productService := NewProductService(productRepositoryMock, logger)

	err := productService.Delete(ctx, productID)
	assert.NoError(t, err)
}
