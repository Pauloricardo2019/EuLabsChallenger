package dto

type ProductPagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

type GetAllProductsResponse struct {
	Products   []Product         `json:"products"`
	Pagination ProductPagination `json:"pagination"`
}

type Product struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func (g *GetAllProductsResponse) ParseFromProductVO(products []Product, limit, offset int, total int64) {
	g.Pagination.Limit = limit
	g.Pagination.Offset = offset
	g.Pagination.Total = total

	for _, product := range products {
		g.Products = append(g.Products, Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
}