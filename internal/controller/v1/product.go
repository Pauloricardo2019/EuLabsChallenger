package v1

import (
	"context"
	"errors"
	"eulabs_challenger/internal/dto"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type productController struct {
	productFacade productFacade
	logger        *zap.Logger
}

func NewControllerProduct(productFacade productFacade, logger *zap.Logger) *productController {
	return &productController{
		productFacade: productFacade,
		logger:        logger,
	}
}

// @Summary create product router
// @Description create product router
// @Tags Product
// @Accept json
// @Param createProductRequest body dto.CreateProductRequest true "create product"
// @Produce json
// @Success 201 {object} dto.CreateProductResponse
// @Failure 500 {object} error
// @Router /eulabs/v1/product [post]
func (p *productController) CreateProduct(c echo.Context) error {
	p.logger.Info("Controller: Creating product")
	ctx := context.Background()

	createProduct := &dto.CreateProductRequest{}

	if err := c.Bind(createProduct); err != nil {
		p.logger.Error("Error binding request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("CreateProductRequest", zap.Any("createProduct", createProduct))

	product, err := p.productFacade.CreateProduct(ctx, createProduct)
	if err != nil {
		p.logger.Error("Error creating product", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Product response", zap.Any("product", product))
	p.logger.Info("Product created")
	return c.JSON(http.StatusCreated, product)
}

// @Summary get product by id router
// @Description get product by id router
// @Tags Product
// @Accept json
// @Param id path int true "id product"
// @Produce json
// @Success 200 {object} dto.GetByProductIDResponse
// @Failure 500 {object} error
// @Router /eulabs/v1/product/{id} [get]
func (p *productController) GetProductByID(c echo.Context) error {
	p.logger.Info("Controller: Getting product by ID")
	ctx := context.Background()

	paramID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Param", zap.Uint64("paramID", paramID))

	product, err := p.productFacade.GetByIDProduct(ctx, paramID)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("product not found")):
			p.logger.Error("Product not found", zap.Error(err))
			return c.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
		default:
			p.logger.Error("Error getting product", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	p.logger.Debug("Product response", zap.Any("product", product))
	p.logger.Info("Product found")
	return c.JSON(http.StatusOK, product)
}

// @Summary get all products by pagination router
// @Description get all products by pagination router
// @Tags Product
// @Accept json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Produce json
// @Success 200 {object} dto.GetAllProductsResponse
// @Failure 500 {object} error
// @Router /eulabs/v1/product [get]
func (p *productController) GetAllProducts(c echo.Context) error {
	p.logger.Info("Controller: Getting all products")
	ctx := context.Background()

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	p.logger.Debug("Params", zap.Int("limit", limit), zap.Int("offset", offset))

	products, err := p.productFacade.GetAllProducts(ctx, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("Products response", zap.Any("products", products))

	p.logger.Info("Controller: Getting all products")
	return c.JSON(http.StatusOK, products)
}

// @Summary update product router
// @Description update product router
// @Tags Product
// @Accept json
// @Param id path int true "id product"
// @Param updateProductRequest body dto.UpdateProductRequest true "update product"
// @Produce json
// @Success 200 {string} string "Product updated successfully"
// @Failure 500 {object} error
// @Router /eulabs/v1/product/{id} [put]
func (p *productController) UpdateProduct(c echo.Context) error {
	p.logger.Info("Controller: Updating product")
	ctx := context.Background()

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}

	p.logger.Debug("Param", zap.Uint64("productID", productID))

	updateProduct := &dto.UpdateProductRequest{}
	if err = c.Bind(updateProduct); err != nil {
		p.logger.Error("Error binding request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("UpdateProductRequest", zap.Any("updateProduct", updateProduct))
	err = p.productFacade.UpdateProduct(ctx, productID, updateProduct)
	if err != nil {
		p.logger.Error("Error updating product", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Info("Product updated successfully")
	return c.JSON(http.StatusOK, "Product updated successfully")
}

// @Summary delete product router
// @Description delete product router
// @Tags Product
// @Accept json
// @Param id path int true "id product"
// @Produce json
// @Success 200 {string} string "Product deleted successfully"
// @Failure 500 {object} error
// @Router /eulabs/v1/product/{id} [delete]
func (p *productController) DeleteProduct(c echo.Context) error {
	p.logger.Info("Controller: Deleting product")
	ctx := context.Background()

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		p.logger.Error("Error parsing param", zap.Error(err))
		return c.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
	}
	p.logger.Debug("Param", zap.Uint64("productID", productID))
	err = p.productFacade.DeleteProduct(ctx, productID)
	if err != nil {
		p.logger.Error("Error deleting product", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
	}

	p.logger.Info("Product deleted successfully")
	return c.JSON(http.StatusOK, "Product deleted successfully")
}
