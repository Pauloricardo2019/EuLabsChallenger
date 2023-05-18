package facade

type productFacade struct {
	productService productService
}

func NewProductFacade(productService productService) ProductFacade {
	return &productFacade{
		productService: productService,
	}
}

func (p *productFacade) CreateProduct(product *Product) (*Product, error) {

}
