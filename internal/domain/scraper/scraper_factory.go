package scraper

import (
	"fmt"
)

type ScraperFactory struct {
	scrapers map[string]Scraper
}

func NewScraperFactory() *ScraperFactory {
	return &ScraperFactory{scrapers: make(map[string]Scraper)}
}

func (sf *ScraperFactory) Register(key string, scraper Scraper) {
	sf.scrapers[key] = scraper
}

func (sf ScraperFactory) GetScraper(key string) (Scraper, error) {
	scraper := sf.scrapers[key]
	if scraper != nil {
		return scraper, nil
	}

	return nil, fmt.Errorf("scraper for key '%s' not registered", key)
}
