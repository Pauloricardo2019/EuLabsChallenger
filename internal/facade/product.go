package facade

import (
	"context"
	"errors"
	"eulabs_challenger/internal/dto"
	"go.uber.org/zap"
)

type productFacade struct {
	productService productService
	logger         *zap.Logger
}

func NewProductFacade(productService productService, logger *zap.Logger) *productFacade {
	return &productFacade{
		productService: productService,
		logger:         logger,
	}
}

func (p *productFacade) CreateProduct(ctx context.Context, product *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	p.logger.Info("Facade: Creating product")

	productVO := product.ConvertToProductVO()

	p.logger.Debug("ProductVO", zap.Any("product", productVO))

	productVO, err := p.productService.Create(ctx, productVO)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Product was created", zap.Any("product", productVO))

	response := &dto.CreateProductResponse{}
	response.ParseFromProductVO(productVO)

	p.logger.Debug("Product response", zap.Any("response", response))
	p.logger.Info("Facade: Product created")
	return response, nil
}

func (p *productFacade) GetByIDProduct(ctx context.Context, id uint64) (*dto.GetByProductIDResponse, error) {
	p.logger.Info("Facade: Getting product by ID")

	p.logger.Debug("Product ID", zap.Uint64("id", id))

	found, productVO, err := p.productService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("product not found")
	}

	p.logger.Debug("Product was found", zap.Any("product", productVO), zap.Bool("found", found))

	response := &dto.GetByProductIDResponse{}
	response.ParseFromProductVO(productVO)

	p.logger.Debug("Product response", zap.Any("response", response))
	p.logger.Info("Facade: Product found")
	return response, nil
}

func (p *productFacade) GetAllProducts(ctx context.Context, limit, offset int) (*dto.GetAllProductsResponse, error) {
	p.logger.Info("Facade: Getting all products")
	products, err := p.productService.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	p.logger.Debug("Products", zap.Any("products", products))

	count, err := p.productService.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	p.logger.Debug("Count", zap.Int64("count", count))

	productsResponse := &dto.GetAllProductsResponse{}
	productsResponse.ParseFromProductVO(products, limit, offset, count)

	p.logger.Debug("Products response", zap.Any("response", productsResponse))
	p.logger.Info("Facade: Products gotten")

	return productsResponse, nil
}

func (p *productFacade) UpdateProduct(ctx context.Context, productID uint64, product *dto.UpdateProductRequest) error {
	p.logger.Info("Facade: Updating product")
	productVO := product.ConvertToProductVO()
	productVO.ID = productID

	p.logger.Debug("ProductVO", zap.Any("product", productVO))
	p.logger.Debug("Product ID", zap.Uint64("id", productID))

	_, err := p.productService.Update(ctx, productVO)
	p.logger.Warn("Product was updated but I'm not using on moment")
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Product updated")
	return nil
}

func (p *productFacade) DeleteProduct(ctx context.Context, id uint64) error {
	p.logger.Info("Facade: Deleting product")
	err := p.productService.Delete(ctx, id)
	if err != nil {
		return err
	}
	p.logger.Info("Facade: Product deleted")
	return nil
}
