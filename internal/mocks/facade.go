package mocks

import (
	"context"
	"eulabs_challenger/internal/dto"
	"github.com/stretchr/testify/mock"
)

type (
	ProductFacadeMock struct {
		mock.Mock
	}
)

func (p *ProductFacadeMock) CreateProduct(ctx context.Context, product *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	args := p.Called(ctx, product)

	var productReq *dto.CreateProductResponse
	var err error

	if args.Get(0) != nil {
		productReq = args.Get(0).(*dto.CreateProductResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return productReq, err
}

func (p *ProductFacadeMock) GetByIDProduct(ctx context.Context, id uint64) (*dto.GetByProductIDResponse, error) {
	args := p.Called(ctx, id)

	var productReq *dto.GetByProductIDResponse
	var err error

	if args.Get(0) != nil {
		productReq = args.Get(0).(*dto.GetByProductIDResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return productReq, err
}

func (p *ProductFacadeMock) GetAllProducts(ctx context.Context, limit, offset int) (*dto.GetAllProductsResponse, error) {
	args := p.Called(ctx, limit, offset)

	var productReq *dto.GetAllProductsResponse
	var err error

	if args.Get(0) != nil {
		productReq = args.Get(0).(*dto.GetAllProductsResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return productReq, err
}

func (p *ProductFacadeMock) UpdateProduct(ctx context.Context, productID uint64, product *dto.UpdateProductRequest) error {
	args := p.Called(ctx, productID, product)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err

}

func (p *ProductFacadeMock) DeleteProduct(ctx context.Context, id uint64) error {
	args := p.Called(ctx, id)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}
