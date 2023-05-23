package facade

import (
	"context"
	"eulabs_challenger/internal/dto"
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

func TestProductFacade_CreateProduct(t *testing.T) {
	ctx := context.Background()

	productServiceMock := &mocks.ProductServiceMock{}

	productToCreate := &dto.CreateProductRequest{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productCreated := &model.Product{
		ID:          1,
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productServiceMock.On("Create", ctx, mock.Anything).
		Return(
			productCreated,
			nil,
		)

	productFacade := NewProductFacade(productServiceMock, logger)

	product, err := productFacade.CreateProduct(ctx, productToCreate)
	assert.NoError(t, err)
	assert.True(t, product.ID == 1)

}

func TestProductFacade_GetByIDProduct(t *testing.T) {
	ctx := context.Background()

	productServiceMock := &mocks.ProductServiceMock{}

	productID := uint64(1)

	productFound := &model.Product{
		ID:          1,
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productServiceMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			productFound,
			nil,
		)

	productFacade := NewProductFacade(productServiceMock, logger)

	product, err := productFacade.GetByIDProduct(ctx, productID)
	assert.NoError(t, err)
	assert.True(t, product.ID == 1)
	assert.True(t, product.Name == "test_product")
	assert.True(t, product.Description == "test_description")
	assert.True(t, product.Price == 10.0)

}

func TestProductFacade_GetAllProducts(t *testing.T) {
	ctx := context.Background()

	productServiceMock := &mocks.ProductServiceMock{}

	limit := 10
	offset := 0

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

	productServiceMock.On("GetAll", ctx, mock.Anything, mock.Anything).
		Return(
			productsFound,
			nil,
		)

	productServiceMock.On("GetCount", ctx).
		Return(
			int64(2),
			nil,
		)

	productFacade := NewProductFacade(productServiceMock, logger)

	products, err := productFacade.GetAllProducts(ctx, limit, offset)
	assert.NoError(t, err)
	assert.True(t, len(products.Products) == 2)
	assert.True(t, products.Pagination.Total == 2)

}

func TestProductFacade_UpdateProduct(t *testing.T) {
	ctx := context.Background()

	productServiceMock := &mocks.ProductServiceMock{}

	productID := uint64(1)

	productToUpdated := &dto.UpdateProductRequest{
		Name:        "test_product_updated",
		Description: "test_description_updated",
		Price:       10.0,
	}

	productUpdated := &model.Product{
		ID:          1,
		Name:        "test_product_updated",
		Description: "test_description_updated",
		Price:       10.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productServiceMock.On("Update", ctx, mock.Anything).
		Return(
			productUpdated,
			nil,
		)

	productServiceMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			productUpdated,
			nil,
		)

	productFacade := NewProductFacade(productServiceMock, logger)

	err := productFacade.UpdateProduct(ctx, productID, productToUpdated)
	assert.NoError(t, err)

}

func TestProductFacade_DeleteProduct(t *testing.T) {
	ctx := context.Background()

	productServiceMock := &mocks.ProductServiceMock{}

	productID := uint64(1)

	productServiceMock.On("Delete", ctx, mock.Anything).
		Return(
			nil,
		)

	productFacade := NewProductFacade(productServiceMock, logger)

	err := productFacade.DeleteProduct(ctx, productID)
	assert.NoError(t, err)

}
