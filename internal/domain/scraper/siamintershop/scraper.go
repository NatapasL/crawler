package siamintershop

import "log"

type Scraper struct {
	updateSeriesByProductService UpdateSeriesByProductService
	productSearchScraper         ProductSearchScraper
}

func NewScraper(seriesUpdater UpdateSeriesByProductService, productSearchScraper ProductSearchScraper) *Scraper {
	return &Scraper{seriesUpdater, productSearchScraper}
}

func (scraper Scraper) Scrape() {
	products, err := scraper.productSearchScraper.Scrape()
	if err != nil {
		log.Println(err)
		return
	}

	scraper.updateSeriesByProductService.UpdateSeries(products)
}
