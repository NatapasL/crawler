package siamintershop

import (
	"fmt"
	"log"
	"manga-crawler/internal/domain/scraper/siamintershop/categorylist"
	"manga-crawler/internal/domain/scraper/siamintershop/productdetail"
	"manga-crawler/internal/domain/scraper/siamintershop/productsearch"
)

type Scraper struct {
	productSearchScraper         productsearch.ProductSearchScraper
	productSearchResponseService productsearch.ResponseService
	productDetailScraper         productdetail.ProductDetailScraper
	productDetailResponseService productdetail.ResponseService
	categoryListScraper          categorylist.CategoryListScraper
}

type ScraperDependencies struct {
	ProductSearchScraper         productsearch.ProductSearchScraper
	ProductSearchResponseService productsearch.ResponseService
	ProductDetailScraper         productdetail.ProductDetailScraper
	ProductDetailResponseService productdetail.ResponseService
	CategoryListScraper          categorylist.CategoryListScraper
}

func NewScraper(deps ScraperDependencies) *Scraper {
	return &Scraper{
		productSearchScraper:         deps.ProductSearchScraper,
		productSearchResponseService: deps.ProductSearchResponseService,
		productDetailScraper:         deps.ProductDetailScraper,
		productDetailResponseService: deps.ProductDetailResponseService,
		categoryListScraper:          deps.CategoryListScraper,
	}
}

func (scraper Scraper) ScrapeAll() {
	categoryList := scraper.ScrapeCategoryList()

	var responseProducts []productsearch.ProductSearchResponseProduct
	for _, category := range categoryList {
		responseProducts = scraper.ScrapeProductSearch(category.CategoryId)
	}

	for _, responseProduct := range responseProducts {
		// scraper.ScrapeProductDetail(responseProduct.ProductId)
		fmt.Println(responseProduct.ProductName)
	}
}

func (scraper Scraper) ScrapeProductSearch(
	categoryId string,
) []productsearch.ProductSearchResponseProduct {
	responseProducts, err := scraper.productSearchScraper.Scrape(categoryId)
	if err != nil {
		log.Println(err)
		return nil
	}

	scraper.productSearchResponseService.AddSeriesNameMatcherIfNotExists(responseProducts)

	return responseProducts
}

func (scraper Scraper) ScrapeProductDetail(productId string) *productdetail.ProductDetail {
	response, err := scraper.productDetailScraper.Scrape(productId)
	if err != nil {
		log.Println(err)
		return nil
	}

	scraper.productDetailResponseService.AddSeriesNameMatcherIfNotExists(*response)

	return response
}

func (scraper Scraper) ScrapeCategoryList() []categorylist.Category {
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
