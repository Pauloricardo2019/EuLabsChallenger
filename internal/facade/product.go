package facade

import (
	"context"
	"errors"
	"eulabs_challenger/internal/dto"
)

type productFacade struct {
	productService productService
}

func NewProductFacade(productService productService) *productFacade {
	return &productFacade{
		productService: productService,
	}
}

func (p *productFacade) CreateProduct(ctx context.Context, product *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	productVO := product.ConvertToProductVO()
	productVO, err := p.productService.Create(ctx, productVO)
	if err != nil {
		return nil, err
	}
	response := &dto.CreateProductResponse{}
	response.ParseFromProductVO(productVO)
	return response, nil
}

func (p *productFacade) GetByIDProduct(ctx context.Context, id uint64) (*dto.GetByProductIDResponse, error) {
	found, productVO, err := p.productService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("product not found")
	}

	response := &dto.GetByProductIDResponse{}
	response.ParseFromProductVO(productVO)
	return response, nil
}

func (p *productFacade) GetAllProducts(ctx context.Context, limit, offset int) (*dto.GetAllProductsResponse, error) {
	products, err := p.productService.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := p.productService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	productsResponse := &dto.GetAllProductsResponse{}
	productsResponse.ParseFromProductVO(products, limit, offset, count)

	return productsResponse, nil
}

func (p *productFacade) UpdateProduct(ctx context.Context, productID uint64, product *dto.UpdateProductRequest) error {
	productVO := product.ConvertToProductVO()

	productVO.ID = productID

	productVO, err := p.productService.Update(ctx, productVO)
	if err != nil {
		return err
	}
	return nil
}

func (p *productFacade) DeleteProduct(ctx context.Context, id uint64) error {
	err := p.productService.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
