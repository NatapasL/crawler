package productdetail

type ProductDetailScraper struct {
	gateway GetProductDetailGateway
}

type ProductDetailScraperDependencies struct {
	Gateway GetProductDetailGateway
}

func NewProductDetailScraper(deps ProductDetailScraperDependencies) *ProductDetailScraper {
	return &ProductDetailScraper{gateway: deps.Gateway}
}

func (pds ProductDetailScraper) Scrape(productId string) (*ProductDetailResponse, error) {
	response, err := pds.gateway.Request(productId)
	if err != nil {
		return nil, err
	}

	return NewProductDetailResponse(response), nil
}
