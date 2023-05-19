package v1

import (
	"context"
	"errors"
	"eulabs_challenger/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type productController struct {
	productFacade productFacade
}

func NewControllerProduct(productFacade productFacade) *productController {
	return &productController{
		productFacade: productFacade,
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

	ctx := context.Background()

	createProduct := &dto.CreateProductRequest{}

	if err := c.Bind(createProduct); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	product, err := p.productFacade.CreateProduct(ctx, createProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

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
	ctx := context.Background()

	paramID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	product, err := p.productFacade.GetByIDProduct(ctx, paramID)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("product not found")):
			return c.JSON(http.StatusNotFound, err)
		default:
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

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
	ctx := context.Background()

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	products, err := p.productFacade.GetAllProducts(ctx, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

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
	ctx := context.Background()

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updateProduct := &dto.UpdateProductRequest{}

	if err = c.Bind(updateProduct); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = p.productFacade.UpdateProduct(ctx, productID, updateProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

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
	ctx := context.Background()

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = p.productFacade.DeleteProduct(ctx, productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Product deleted successfully")
}
