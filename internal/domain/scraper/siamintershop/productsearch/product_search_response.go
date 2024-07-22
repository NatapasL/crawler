package productsearch

type ProductSearchResponse struct {
	Total    string                         `json:"total"`
	Limit    string                         `json:"limit"`
	Offset   int                            `json:"offset"`
	Products []ProductSearchResponseProduct `json:"products"`
}

type ProductSearchResponseProduct struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
}
