package main

import (
	"eulabs_challenger/adapter/provider"
	"eulabs_challenger/adapter/repository"
	"eulabs_challenger/adapter/rest"
	"eulabs_challenger/internal/config"
	"eulabs_challenger/internal/controller/v1"
	"eulabs_challenger/internal/facade"
	"eulabs_challenger/internal/service"
	"fmt"
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
	//Config
	cfg := config.NewConfig().GetConfig()

	fmt.Println("Setup config")

	//Get Database Connection
	dbProvider := provider.NewDatabaseProvider(cfg)
	dbConn, err := dbProvider.Connect()
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Setup database")

	//Get Repositories
	productRepository := repository.NewProductRepository(dbConn)

	fmt.Println("Setup repositories")

	//Get Services
	productService := service.NewProductService(productRepository)

	fmt.Println("Setup services")

	//Get Facades
	productFacade := facade.NewProductFacade(productService)

	fmt.Println("Setup facades")

	//Get Controllers
	productController := v1.NewControllerProduct(productFacade)
	healthCheckController := v1.NewHealthCheckController()

	fmt.Println("Setup controllers")

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			HealthCheckController: healthCheckController,
			ProductController:     productController,
		},
	)

	fmt.Println("Run server...")
	//swag init --parseDependency -g ./cmd/api/main.go
	serverRest.StartListening()

}
