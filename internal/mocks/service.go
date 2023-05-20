package mocks

import (
	"context"
	"eulabs_challenger/internal/model"
	"github.com/stretchr/testify/mock"
)

type (
	productServiceMock struct {
		mock.Mock
	}
)

func (p *productServiceMock) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	args := p.Called(ctx, product)

	var productReq *model.Product
	var err error

	if args.Get(0) != nil {
		productReq = args.Get(0).(*model.Product)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return productReq, err
}

func (p *productServiceMock) GetCount(ctx context.Context) (int64, error) {
	args := p.Called(ctx)

	var err error
	var count int64

	if args.Get(0) != nil {
		count = args.Get(0).(int64)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return count, err
}

func (p *productServiceMock) GetByID(ctx context.Context, id uint64) (bool, *model.Product, error) {
	args := p.Called(ctx, id)

	found := args.Get(0).(bool)

	var product *model.Product
	var err error

	if args.Get(1) != nil {
		product = args.Get(1).(*model.Product)
	}

	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}

	return found, product, err
}

func (p *productServiceMock) GetAll(ctx context.Context, limit, offset int) ([]model.Product, error) {
	args := p.Called(ctx, limit, offset)

	var products []model.Product
	var err error

	if args.Get(0) != nil {
		products = args.Get(0).([]model.Product)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return products, err
}

func (p *productServiceMock) Update(ctx context.Context, product *model.Product) (*model.Product, error) {
	args := p.Called(ctx, product)

	var productReq *model.Product
	var err error

	if args.Get(0) != nil {
		productReq = args.Get(0).(*model.Product)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return productReq, err
}

func (p *productServiceMock) Delete(ctx context.Context, id uint64) error {
	args := p.Called(ctx, id)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}
