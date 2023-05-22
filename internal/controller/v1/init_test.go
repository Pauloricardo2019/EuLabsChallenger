package v1_test

import (
	"eulabs_challenger/adapter/rest"
	"eulabs_challenger/internal/controller/v1"
	facadeMocks "eulabs_challenger/internal/mocks"
	"eulabs_challenger/internal/model"
	"go.uber.org/zap"
	"testing"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

type Facade struct {
	ProductControllerMock *facadeMocks.ProductFacadeMock
}

func setupTestRouter(t *testing.T) (*rest.ServerRest, Facade) {
	t.Helper()

	facades := Facade{
		ProductControllerMock: &facadeMocks.ProductFacadeMock{},
	}

	cfg := &model.Config{}

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			ProductController:     v1.NewControllerProduct(facades.ProductControllerMock, logger),
			HealthCheckController: v1.NewHealthCheckController(),
		},
	)

	return serverRest, facades

}
