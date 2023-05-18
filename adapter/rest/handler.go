package rest

import (
	"eulabs_challenger/internal/model"
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
	s.Engine.GET("/health", s.controllers.HealthCheckController.HealthCheck)

}
