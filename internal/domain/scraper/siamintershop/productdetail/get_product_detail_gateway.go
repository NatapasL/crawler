package productdetail

type GetProductDetailGateway interface {
	Request(productId string) (*ProductDetail, error)
}
