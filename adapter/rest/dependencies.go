package rest

import "github.com/labstack/echo/v4"

type healthCheckController interface {
	HealthCheck()
}

type productController interface {
	CreateProduct(c *echo.Context)
	GetProductByID(c *echo.Context)
	GetAllProducts(c *echo.Context)
	UpdateProduct(c *echo.Context)
	DeleteProduct(c *echo.Context)
}