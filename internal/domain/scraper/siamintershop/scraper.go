package siamintershop

import (
	siamintershopproductdetail "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_detail"
	siamintershopproductsearch "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_search"
)

type Scraper struct {
	productSearchScraper         siamintershopproductsearch.ProductSearchScraper
	productSearchResponseService siamintershopproductsearch.ResponseService

	productDetailScraper siamintershopproductdetail.ProductDetailScraper
}

func NewScraper(
	productSearchScraper siamintershopproductsearch.ProductSearchScraper,
	productSearchResponseService siamintershopproductsearch.ResponseService,
	productDetailScraper siamintershopproductdetail.ProductDetailScraper,
) *Scraper {
	return &Scraper{productSearchScraper, productSearchResponseService, productDetailScraper}
}

func (scraper Scraper) Scrape() {
	// responseProducts, err := scraper.productSearchScraper.Scrape()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// product := model.New

	// scraper.updateSeriesByProductService.UpdateSeries(products)

	scraper.productDetailScraper.Scrape("11000408800022331")
}
