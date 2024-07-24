package siamintershop

import (
	"errors"
	"fmt"
	"log"
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/catalogupdate/wholesale"
	"manga-crawler/internal/domain/scraper"
	"manga-crawler/internal/domain/scraper/siamintershop/categorylist"
	"manga-crawler/internal/domain/scraper/siamintershop/productdetail"
	"manga-crawler/internal/domain/scraper/siamintershop/productsearch"

	"github.com/google/uuid"
)

const SiamintershopID = "22e26f1c-a8b2-4728-ab6d-7045798afc1a"

var sicPublisherID = uuid.Must(uuid.Parse("01eb250e-57ad-4be1-8906-dc1527de6238"))

type Scraper struct {
	productSearchScraper productsearch.ProductSearchScraper
	productDetailScraper productdetail.ProductDetailScraper
	categoryListScraper  categorylist.CategoryListScraper
	nameCleaner          scraper.NameCleaner
	seriesFinder         series.SeriesNameFinder
}

type ScraperDependencies struct {
	ProductSearchScraper productsearch.ProductSearchScraper
	ProductDetailScraper productdetail.ProductDetailScraper
	CategoryListScraper  categorylist.CategoryListScraper
	NameCleaner          scraper.NameCleaner
	SeriesFinder         series.SeriesNameFinder
}

func NewScraper(deps ScraperDependencies) *Scraper {
	return &Scraper{
		productSearchScraper: deps.ProductSearchScraper,
		productDetailScraper: deps.ProductDetailScraper,
		categoryListScraper:  deps.CategoryListScraper,
		nameCleaner:          deps.NameCleaner,
		seriesFinder:         deps.SeriesFinder,
	}
}

func (scraper Scraper) Scrape(next <-chan bool, product chan<- wholesale.Product) error {
	categoryList := scraper.ScrapeCategoryList()

	var errs []error
	for _, category := range categoryList {
		responseProducts := scraper.ScrapeProductSearch(category.CategoryId)
		for _, responseProduct := range responseProducts {
			<-next
			p, err := responseProduct.ToProduct(scraper.seriesFinder, scraper.nameCleaner, sicPublisherID)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			product <- p
		}
	}
	close(product)

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
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
