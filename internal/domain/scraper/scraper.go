package scraper

import (
	"manga-crawler/internal/domain/catalogupdate/product"
)

type Scraper interface {
	Scrape(chan<- product.Product) error
}
