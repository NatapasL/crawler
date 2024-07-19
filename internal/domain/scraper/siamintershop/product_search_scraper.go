package siamintershop

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
	config               productSearchScraperConfig
	intervalConfig       productSearchScraperIntervalConfig
	productSearchFetcher ProductSearchFetcher
}

func NewProductSearchScraper(productSearchFetcher ProductSearchFetcher) *ProductSearchScraper {
	return &ProductSearchScraper{
		productSearchFetcher: productSearchFetcher,
		config: productSearchScraperConfig{
			categoryId: "654",
		},
		intervalConfig: productSearchScraperIntervalConfig{
			chunkSize: 59,
			delay:     1 * time.Second,
		},
	}
}

func (pss ProductSearchScraper) Scrape() ([]SiamintershopProduct, error) {
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
	response, err := pss.productSearchFetcher.Fetch(productSearchFilter{
		Limit:             1,
		Offset:            0,
		CategoryId:        pss.config.categoryId,
		CategoryWithChild: true,
	})

	siamintershopResponse := NewSiamintershopResponseFromBytes(response)
	if err != nil {
		return 0, err
	}

	total, err := strconv.Atoi(siamintershopResponse.Total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (pss ProductSearchScraper) iterateGetProducts(total int) ([]SiamintershopProduct, error) {
	var products []SiamintershopProduct
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

func (pss ProductSearchScraper) getProducts(limit int, offset int) ([]SiamintershopProduct, error) {
	response, err := pss.productSearchFetcher.Fetch(productSearchFilter{
		Limit:             limit,
		Offset:            offset,
		CategoryId:        pss.config.categoryId,
		CategoryWithChild: true,
	})

	siamintershopResponse := NewSiamintershopResponseFromBytes(response)
	if siamintershopResponse == nil {
		return nil, err
	}

	return siamintershopResponse.Products, nil
}
