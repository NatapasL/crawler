package main

import (
	"manga-crawler/internal/domain/scraper/namecleaner"
	"manga-crawler/internal/domain/scraper/siamintershop"
	postgresrepository "manga-crawler/internal/infrastructure/repository/postgres_repository"
)

type siamintershopScraperContainer struct {
	scraper *siamintershop.Scraper
}

func NewSiamintershopScraperContainer(
	seriesRepository *postgresrepository.SeriesRepository,
	seriesNameMatcherRepository *postgresrepository.SeriesNameMatcherRepository,
) *siamintershopScraperContainer {
	regexpPattern := namecleaner.NewRegexpPattern(siamintershop.GetNameCleanerPattern())
	nameCleaner := namecleaner.NewNameCleaner(regexpPattern)
	updateSeriesByProduct := siamintershop.NewUpdateSeriesByProductService(nameCleaner, seriesRepository, seriesNameMatcherRepository)
	productSearchFetcher := siamintershop.NewProductSearchFetcher()
	productSearchScraper := siamintershop.NewProductSearchScraper(*productSearchFetcher)
	scraper := siamintershop.NewScraper(*updateSeriesByProduct, *productSearchScraper)

	return &siamintershopScraperContainer{scraper}
}
