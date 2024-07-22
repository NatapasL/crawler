package subproduct

type GetSubProductsGateway interface {
	Request(productId string) ([]byte, error)
}
