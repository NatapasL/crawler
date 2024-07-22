package productdetail

type GetProductDetailGateway interface {
	Request(productId string) ([]byte, error)
}
