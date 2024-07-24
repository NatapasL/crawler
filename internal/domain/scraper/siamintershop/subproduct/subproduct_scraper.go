package subproduct

type SubProductScraper struct {
	gateway GetSubProductsGateway
}

type SubProductScraperDependencies struct {
	Gateway GetSubProductsGateway
}

func NewSubProductScraper(deps SubProductScraperDependencies) *SubProductScraper {
	return &SubProductScraper{gateway: deps.Gateway}
}

func (pds SubProductScraper) Scrape(productId string) ([]SubProduct, error) {
	productDetail, err := pds.gateway.Request(productId)
	if err != nil {
		return nil, err
	}

	return productDetail, nil
}
