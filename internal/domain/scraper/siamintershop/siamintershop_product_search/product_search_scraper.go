package siamintershopproductsearch

import (
	"strconv"
	"time"
)

type productSearchScraperIntervalConfig struct {
	chunkSize int
	delay     time.Duration
}

type productSearchScraperConfig struct {
	categoryId string
}

type ProductSearchScraper struct {
	config         productSearchScraperConfig
	intervalConfig productSearchScraperIntervalConfig
	gateway        ProductSearchGateway
}

func NewProductSearchScraper(productSearchGateway ProductSearchGateway) *ProductSearchScraper {
	return &ProductSearchScraper{
		gateway: productSearchGateway,
		config: productSearchScraperConfig{
			categoryId: "654",
		},
		intervalConfig: productSearchScraperIntervalConfig{
			chunkSize: 59,
			delay:     1 * time.Second,
		},
	}
}

func (pss ProductSearchScraper) Scrape() ([]ProductSearchResponseProduct, error) {
	total, err := pss.getTotalProducts()
	if err != nil {
		return nil, err
	}

	products, err := pss.iterateGetProducts(total)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pss ProductSearchScraper) getTotalProducts() (int, error) {
	response, err := pss.gateway.Request(pss.config.categoryId, 0, 1)

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

func (pss ProductSearchScraper) iterateGetProducts(total int) ([]ProductSearchResponseProduct, error) {
	var products []ProductSearchResponseProduct
	chunkSize := pss.intervalConfig.chunkSize

	for i := 0; i*chunkSize < total; i++ {
		time.Sleep(pss.intervalConfig.delay)

		responseProducts, err := pss.getProducts(chunkSize, i*chunkSize)
		if err != nil {
			return nil, err
		}
		products = append(products, responseProducts...)
	}

	return products, nil
}

func (pss ProductSearchScraper) getProducts(limit int, offset int) ([]ProductSearchResponseProduct, error) {
	response, err := pss.gateway.Request(pss.config.categoryId, offset, limit)

	siamintershopResponse := NewProductSearchResponse(response)
	if siamintershopResponse == nil {
		return nil, err
	}

	return siamintershopResponse.Products, nil
}
