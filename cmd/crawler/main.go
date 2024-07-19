package main

import (
	"log"
	"manga-crawler/config"
	"manga-crawler/internal/domain/scraper/namecleaner"
	"manga-crawler/internal/domain/scraper/siamintershop"
	"manga-crawler/internal/infrastructure/database/postgres"
)

func main() {
	conf := config.GetConfig()
	db := postgres.GetConnection(conf.Postgres)
	if db == nil {
		log.Println("Cant connect to db...")
		return
	}

	seriesRepository := postgres.NewSeriesRepository(db)

	regexpPattern := namecleaner.NewRegexpPattern(siamintershop.GetNameCleanerPattern())
	nameCleaner := namecleaner.NewNameCleaner(regexpPattern)
	seriesUpdater := siamintershop.NewSeriesUpdater(nameCleaner, seriesRepository)
	productSearchFetcher := siamintershop.NewProductSearchFetcher()
	productSearchScraper := siamintershop.NewProductSearchScraper(*productSearchFetcher)
	scraper := siamintershop.NewScraper(*seriesUpdater, *productSearchScraper)

	scraper.Scrape()
}
