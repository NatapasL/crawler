package siamintershop

import (
	"fmt"
	"log"
	"manga-crawler/internal/domain/catalogupdate/wholesale"
	"manga-crawler/internal/domain/scraper"
	"manga-crawler/internal/domain/scraper/siamintershop/categorylist"
	"manga-crawler/internal/domain/scraper/siamintershop/productdetail"
	"manga-crawler/internal/domain/scraper/siamintershop/productsearch"
)

const SiamintershopID = "22e26f1c-a8b2-4728-ab6d-7045798afc1a"

type Scraper struct {
	productSearchScraper productsearch.ProductSearchScraper
	productDetailScraper productdetail.ProductDetailScraper
	categoryListScraper  categorylist.CategoryListScraper
	nameCleaner          scraper.NameCleaner
}

type ScraperDependencies struct {
	ProductSearchScraper productsearch.ProductSearchScraper
	ProductDetailScraper productdetail.ProductDetailScraper
	CategoryListScraper  categorylist.CategoryListScraper
	NameCleaner          scraper.NameCleaner
}

func NewScraper(deps ScraperDependencies) *Scraper {
	return &Scraper{
		productSearchScraper: deps.ProductSearchScraper,
		productDetailScraper: deps.ProductDetailScraper,
		categoryListScraper:  deps.CategoryListScraper,
		nameCleaner:          deps.NameCleaner,
	}
}

func (scraper Scraper) Scrape() ([]wholesale.Product, error) {
	return make([]wholesale.Product, 0), nil
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

	return responseProducts
}

func (scraper Scraper) ScrapeProductDetail(productId string) *productdetail.ProductDetail {
	response, err := scraper.productDetailScraper.Scrape(productId)
	if err != nil {
		log.Println(err)
		return nil
	}

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
