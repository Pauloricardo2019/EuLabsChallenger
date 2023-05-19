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

func (p *productController) CreateProduct(c echo.Context) {

	ctx := context.Background()

	createProduct := &dto.CreateProductRequest{}

	if err := c.Bind(createProduct); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := p.productFacade.CreateProduct(ctx, createProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, product)

}

func (p *productController) GetProductByID(c echo.Context) {
	ctx := context.Background()

	paramID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := p.productFacade.GetByIDProduct(ctx, paramID)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("product not found")):
			c.JSON(http.StatusNotFound, err)
		default:
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, product)

}

func (p *productController) GetAllProducts(c echo.Context) {
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p *productController) UpdateProduct(c echo.Context) {
	ctx := context.Background()

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updateProduct := &dto.UpdateProductRequest{}

	if err = c.Bind(updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.productFacade.UpdateProduct(ctx, productID, updateProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Product updated successfully")

}

func (p *productController) DeleteProduct(c echo.Context) {
	ctx := context.Background()

	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = p.productFacade.DeleteProduct(ctx, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Product deleted successfully")
}