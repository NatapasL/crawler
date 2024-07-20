package siamintershopproductdetail

type ProductDetailScraper struct {
	gateway GetProductDetailGateway
}

func NewProductDetailScraper(gateway GetProductDetailGateway) *ProductDetailScraper {
	return &ProductDetailScraper{gateway}
}

func (pds ProductDetailScraper) Scrape(productId string) (*ProductDetailResponse, error) {
	response, err := pds.gateway.Request(productId)
	if err != nil {
		return nil, err
	}

	return NewProductDetailResponse(response), nil
}
