package siamintershop

import (
	"log"
	siamintershopcategorylist "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_category_list"
	siamintershopproductdetail "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_detail"
	siamintershopproductsearch "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_search"
)

type Scraper struct {
	productSearchScraper         siamintershopproductsearch.ProductSearchScraper
	productSearchResponseService siamintershopproductsearch.ResponseService

	productDetailScraper         siamintershopproductdetail.ProductDetailScraper
	productDetailResponseService siamintershopproductdetail.ResponseService

	categoryListScraper siamintershopcategorylist.CategoryListScraper
}

func NewScraper(
	productSearchScraper siamintershopproductsearch.ProductSearchScraper,
	productSearchResponseService siamintershopproductsearch.ResponseService,

	productDetailScraper siamintershopproductdetail.ProductDetailScraper,
	productDetailResponseService siamintershopproductdetail.ResponseService,

	categoryListScraper siamintershopcategorylist.CategoryListScraper,
) *Scraper {
	return &Scraper{
		productSearchScraper,
		productSearchResponseService,
		productDetailScraper,
		productDetailResponseService,
		categoryListScraper,
	}
}

func (scraper Scraper) ScrapeProductSearch() []siamintershopproductsearch.ProductSearchResponseProduct {
	responseProducts, err := scraper.productSearchScraper.Scrape()
	if err != nil {
		log.Println(err)
		return nil
	}

	scraper.productSearchResponseService.AddSeriesNameMatcherIfNotExists(responseProducts)

	return responseProducts
}

func (scraper Scraper) ScrapeProductDetail() *siamintershopproductdetail.ProductDetailResponse {
	response, err := scraper.productDetailScraper.Scrape("11000408800022331")
	if err != nil {
		log.Println(err)
		return nil
	}

	scraper.productDetailResponseService.AddSeriesNameMatcherIfNotExists(*response)

	return response
}

func (scraper Scraper) ScrapeCategoryList() []siamintershopcategorylist.CategoryResponse {
	response, err := scraper.categoryListScraper.Scrape()
	if err != nil {
		log.Println(err)
		return nil
	}

	return response
}

func (scraper Scraper) ScrapeProductSubproducts() {

}
