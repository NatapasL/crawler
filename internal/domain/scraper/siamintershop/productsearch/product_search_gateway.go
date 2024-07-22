package productsearch

type ProductSearchGateway interface {
	Request(categoryId string, offset int, limit int) (*ProductSearchResponse, error)
}
