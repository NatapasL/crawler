package siamintershopproductsearch

import (
	"strconv"
	"time"
)

type ProductSearchScraper struct {
	intervalConfig productSearchScraperIntervalConfig
	gateway        ProductSearchGateway
}

type ProductSearchScraperDependencies struct {
	Gateway ProductSearchGateway
}

func NewProductSearchScraper(deps ProductSearchScraperDependencies) *ProductSearchScraper {
	return &ProductSearchScraper{
		gateway:        deps.Gateway,
		intervalConfig: defaultIntervalConfig(),
	}
}

func (pss ProductSearchScraper) Scrape(categoryId string) ([]ProductSearchResponseProduct, error) {
	total, err := pss.getTotalProducts(categoryId)
	if err != nil {
		return nil, err
	}

	products, err := pss.iterateGetProducts(total, categoryId)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pss ProductSearchScraper) getTotalProducts(categoryId string) (int, error) {
	response, err := pss.gateway.Request(categoryId, 0, 1)

	siamintershopResponse := NewProductSearchResponse(response)
	if err != nil {
		return 0, err
	}

	total, err := strconv.Atoi(siamintershopResponse.Total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (pss ProductSearchScraper) iterateGetProducts(total int, categoryId string) ([]ProductSearchResponseProduct, error) {
	var products []ProductSearchResponseProduct
	chunkSize := pss.intervalConfig.chunkSize

	for i := 0; i*chunkSize < total; i++ {
		time.Sleep(pss.intervalConfig.delay)

		responseProducts, err := pss.getProducts(categoryId, chunkSize, i*chunkSize)
		if err != nil {
			return nil, err
		}
		products = append(products, responseProducts...)
	}

	return products, nil
}

func (pss ProductSearchScraper) getProducts(categoryId string, limit int, offset int) ([]ProductSearchResponseProduct, error) {
	response, err := pss.gateway.Request(categoryId, offset, limit)

	siamintershopResponse := NewProductSearchResponse(response)
	if siamintershopResponse == nil {
		return nil, err
	}

	return siamintershopResponse.Products, nil
}
