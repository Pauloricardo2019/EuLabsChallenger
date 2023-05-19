package api

import (
	"eulabs_challenger/adapter/provider"
	"eulabs_challenger/adapter/repository"
	"eulabs_challenger/adapter/rest"
	"eulabs_challenger/internal/config"
	"eulabs_challenger/internal/controller/v1"
	"eulabs_challenger/internal/facade"
	"eulabs_challenger/internal/service"
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
	cfg := config.NewConfig()

	//Get Database Connection
	dbProvider := provider.NewDatabaseProvider(cfg.GetConfig())

	dbConn, err := dbProvider.Connect()
	if err != nil {
		panic(err)
		return
	}

	//Get Repositories
	productRepository := repository.NewProductRepository(dbConn)

	//Get Services
	productService := service.NewProductService(productRepository)

	//Get Facades
	productFacade := facade.NewProductFacade(productService)

	//Get Controllers
	productController := v1.NewControllerProduct(productFacade)

	serverRest := rest.NewRestServer(
		cfg.GetConfig(),
		&rest.Controllers{
			ProductController: productController,
		},
	)

	serverRest.StartListening()

}
