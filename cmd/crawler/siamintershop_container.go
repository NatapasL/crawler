package main

import (
	"manga-crawler/internal/domain/scraper/namecleaner"
	"manga-crawler/internal/domain/scraper/siamintershop"
	siamintershopgapi "manga-crawler/internal/infrastructure/api_gateway/siamintershop_api"
	postgresrepository "manga-crawler/internal/infrastructure/repository/postgres_repository"
)

type siamintershopScraperContainer struct {
	scraper *siamintershop.Scraper
}

func NewSiamintershopScraperContainer(
	seriesRepository *postgresrepository.SeriesRepository,
	seriesNameMatcherRepository *postgresrepository.SeriesNameMatcherRepository,
) *siamintershopScraperContainer {
	// api gateway
	apiRequest := siamintershopgapi.ApiRequest{}
	productSearchApi := siamintershopgapi.NewProductSearchApi(apiRequest)

	// app
	regexpPattern := namecleaner.NewRegexpPattern(siamintershop.GetNameCleanerPattern())
	nameCleaner := namecleaner.NewNameCleaner(regexpPattern)
	updateSeriesByProduct := siamintershop.NewUpdateSeriesByProductService(nameCleaner, seriesRepository, seriesNameMatcherRepository)
	productSearchScraper := siamintershop.NewProductSearchScraper(*productSearchApi)
	scraper := siamintershop.NewScraper(*updateSeriesByProduct, *productSearchScraper)

	return &siamintershopScraperContainer{scraper}
}
