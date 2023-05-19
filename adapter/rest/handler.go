package rest

import (
	"eulabs_challenger/internal/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Controllers struct {
	HealthCheckController healthCheckController
	ProductController     productController
}

type ServerRest struct {
	httpServer  *http.Server
	Engine      *echo.Echo
	config      *model.Config
	controllers *Controllers
}

func NewRestServer(cfg *model.Config, controllers *Controllers) *ServerRest {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	server := &ServerRest{
		Engine:      e,
		config:      cfg,
		controllers: controllers,
	}

	server.registerRoutes()
	return server
}

func (s *ServerRest) registerRoutes() {
	//routeV1 := s.Engine.Group("/v1")
	//{
	//	routeV1.GET("/health", s.controllers.HealthCheckController.HealthCheck)
	//
	//	productGroup := routeV1.Group("/product")
	//	{
	//		productGroup.POST("", s.controllers.ProductController.CreateProduct)
	//		productGroup.GET("/:id", s.controllers.ProductController.GetProductByID)
	//		productGroup.GET("", s.controllers.ProductController.GetAllProducts)
	//		productGroup.PUT("/:id", s.controllers.ProductController.UpdateProduct)
	//		productGroup.DELETE("/:id", s.controllers.ProductController.DeleteProduct)
	//	}
	//
	//}

}

func (s *ServerRest) StartListening() {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.RestPort),
		Handler: s.Engine,
	}

	fmt.Println("Listening on port", s.config.RestPort)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err.Error())
	}
}
