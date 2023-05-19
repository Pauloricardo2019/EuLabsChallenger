package rest

import "github.com/labstack/echo/v4"

type (
	healthCheckController interface {
		HealthCheck(c echo.Context)
	}

	productController interface {
		CreateProduct(c echo.Context)
		GetProductByID(c echo.Context)
		GetAllProducts(c echo.Context)
		UpdateProduct(c echo.Context)
		DeleteProduct(c echo.Context)
	}
)
