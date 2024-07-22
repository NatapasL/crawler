package siamintershop

import (
	"manga-crawler/cmd/crawler/container/seriesnamematcher"
	"manga-crawler/internal/domain/scraper/siamintershop/productsearch"
)

type siamintershopProductSearchModule struct {
	ProductSearchScraper *productsearch.ProductSearchScraper
	ResponseService      *productsearch.ResponseService
}

type siamintershopProductSearchModuleDependencies struct {
	ApiModule               siamintershopApiModule
	NameCleanerModule       siamintershopNameCleanerModule
	SeriesNameMatcherModule seriesnamematcher.SeriesNameMatcherModule
}

func initializeSiamintershopProductSearchModule(
	deps siamintershopProductSearchModuleDependencies,
) siamintershopProductSearchModule {
	scraper := productsearch.NewProductSearchScraper(
		productsearch.ProductSearchScraperDependencies{Gateway: deps.ApiModule.ProductSearchApi},
	)
	responseService := productsearch.NewResponseService(
		productsearch.ResponseServiceDependencies{
			NameCleaner:              *deps.NameCleanerModule.NameCleaner,
			SeriesNameMatcherService: *deps.SeriesNameMatcherModule.SeriesNameMatcherService,
		},
	)

	return siamintershopProductSearchModule{
		ProductSearchScraper: scraper,
		ResponseService:      responseService,
	}
}
