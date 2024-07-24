package scraper

import (
	"fmt"

	"github.com/google/uuid"
)

type ScraperFactory struct {
	scrapers map[uuid.UUID]Scraper
}

func NewScraperFactory() *ScraperFactory {
	return &ScraperFactory{scrapers: make(map[uuid.UUID]Scraper)}
}

func (sf *ScraperFactory) Register(wholesaleID uuid.UUID, scraper Scraper) {
	sf.scrapers[wholesaleID] = scraper
}

func (sf ScraperFactory) GetScraper(wholesaleID uuid.UUID) (Scraper, error) {
	scraper := sf.scrapers[wholesaleID]
	if scraper != nil {
		return scraper, nil
	}

	return nil, fmt.Errorf("scraper for wholesale %s not registered", wholesaleID.String())
}
