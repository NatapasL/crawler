package subproduct

type GetSubProductsGateway interface {
	Request(productId string) ([]SubProduct, error)
}
