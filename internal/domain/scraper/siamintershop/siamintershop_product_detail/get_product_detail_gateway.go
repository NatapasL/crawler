package siamintershopproductdetail

type GetProductDetailGateway interface {
	Request(productId string) ([]byte, error)
}
