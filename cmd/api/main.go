package main

import (
	"eulabs_challenger/adapter/provider"
	"eulabs_challenger/adapter/repository"
	"eulabs_challenger/adapter/rest"
	"eulabs_challenger/internal/config"
	"eulabs_challenger/internal/controller/v1"
	"eulabs_challenger/internal/facade"
	"eulabs_challenger/internal/service"
	"go.uber.org/zap"
	"time"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	logger := zap.NewExample()

	defer logger.Sync()

	//Config
	cfg := config.NewConfig(logger).GetConfig()

	logger.Info("Setup config",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Database Connection
	dbProvider := provider.NewDatabaseProvider(cfg, logger)
	dbConn, err := dbProvider.Connect()
	if err != nil {
		panic(err)
		return
	}

	logger.Info("Setup database",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Repositories
	productRepository := repository.NewProductRepository(dbConn, logger)

	logger.Info("Setup repositories",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Services
	productService := service.NewProductService(productRepository, logger)

	logger.Info("Setup services",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Facades
	productFacade := facade.NewProductFacade(productService, logger)

	logger.Info("Setup facades",
		zap.Time("StartedAt", time.Now()),
	)

	//Get Controllers
	productController := v1.NewControllerProduct(productFacade, logger)
	healthCheckController := v1.NewHealthCheckController()

	logger.Info("Setup controllers",
		zap.Time("StartedAt", time.Now()),
	)

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			HealthCheckController: healthCheckController,
			ProductController:     productController,
		},
	)

	logger.Info("Setup server",
		zap.Time("StartedAt", time.Now()),
	)

	serverRest.StartListening()

}
