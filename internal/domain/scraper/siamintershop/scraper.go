package siamintershop

import (
	"fmt"
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

func (scraper Scraper) ScrapeAll() {
	categoryList := scraper.ScrapeCategoryList()

	var responseProducts []siamintershopproductsearch.ProductSearchResponseProduct
	for _, category := range categoryList {
		responseProducts = scraper.ScrapeProductSearch(category.CategoryId)
	}

	for _, responseProduct := range responseProducts {
		// scraper.ScrapeProductDetail(responseProduct.ProductId)
		fmt.Println(responseProduct.ProductName)
	}
}

func (scraper Scraper) ScrapeProductSearch(categoryId string) []siamintershopproductsearch.ProductSearchResponseProduct {
	responseProducts, err := scraper.productSearchScraper.Scrape(categoryId)
	if err != nil {
		log.Println(err)
		return nil
	}

	scraper.productSearchResponseService.AddSeriesNameMatcherIfNotExists(responseProducts)

	return responseProducts
}

func (scraper Scraper) ScrapeProductDetail(productId string) *siamintershopproductdetail.ProductDetailResponse {
	response, err := scraper.productDetailScraper.Scrape(productId)
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

	fmt.Printf("%+v", response)

	return response
}

func (scraper Scraper) ScrapeProductSubproducts() {

}
