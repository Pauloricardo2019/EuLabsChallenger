package service

import (
	"context"
	"eulabs_challenger/internal/mocks"
	"eulabs_challenger/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProductService_Create(t *testing.T) {
	ctx := context.Background()

	productServiceMock := &mocks.ProductRepositoryMock{}

	productMock := &model.Product{
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productMockCreated := &model.Product{
		ID:          1,
		Name:        "test_product",
		Description: "test_description",
		Price:       10.0,
	}

	productServiceMock.On("Create", ctx, mock.Anything).
		Return(
			productMockCreated,
			nil,
		)

	productService := NewProductService(productServiceMock)

	productCreated, err := productService.Create(ctx, productMock)
	assert.NoError(t, err)
	assert.True(t, productCreated.ID == 1)

}
