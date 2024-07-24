package scraper

import "manga-crawler/internal/domain/catalogupdate/wholesale"

type Scraper interface {
	Scrape() ([]wholesale.Product, error)
}
