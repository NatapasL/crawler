package scraper

import "manga-crawler/internal/domain/catalogupdate/wholesale"

type Scraper interface {
	Scrape(next <-chan bool, product chan<- wholesale.Product) error
}
