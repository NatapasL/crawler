package siamintershopproductdetail

import (
	"fmt"
	siamintershopgateway "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_gateway"
)

type ProductDetailScraper struct {
	gateway siamintershopgateway.GetProductDetailGateway
}

func NewProductDetailScraper(gateway siamintershopgateway.GetProductDetailGateway) *ProductDetailScraper {
	return &ProductDetailScraper{gateway}
}

func (pds ProductDetailScraper) Scrape(productId string) {
	response, err := pds.gateway.Request(productId)
	if err != nil {
		return
	}

	product := NewProductDetailResponse(response)

	fmt.Printf("%+v", product)
}
