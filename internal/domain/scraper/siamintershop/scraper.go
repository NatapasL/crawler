package siamintershop

import "log"

type Scraper struct {
	seriesUpdater        SeriesUpdater
	productSearchScraper ProductSearchScraper
}

func NewScraper(seriesUpdater SeriesUpdater, productSearchScraper ProductSearchScraper) *Scraper {
	return &Scraper{seriesUpdater, productSearchScraper}
}

func (scraper Scraper) Scrape() {
	products, err := scraper.productSearchScraper.Scrape()
	if err != nil {
		log.Println(err)
		return
	}

	scraper.seriesUpdater.UpdateSeries(products)
}
