package container

import siamintershopproductsearch "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_search"

type siamintershopProductSearchModule struct {
	ProductSearchScraper *siamintershopproductsearch.ProductSearchScraper
	ResponseService      *siamintershopproductsearch.ResponseService
}

type siamintershopProductSearchModuleDependencies struct {
	ApiModule               siamintershopApiModule
	NameCleanerModule       siamintershopNameCleanerModule
	SeriesNameMatcherModule SeriesNameMatcherModule
}

func initializeSiamintershopProductSearchModule(
	deps siamintershopProductSearchModuleDependencies,
) siamintershopProductSearchModule {
	scraper := siamintershopproductsearch.NewProductSearchScraper(
		siamintershopproductsearch.ProductSearchScraperDependencies{Gateway: deps.ApiModule.ProductSearchApi},
	)
	responseService := siamintershopproductsearch.NewResponseService(
		siamintershopproductsearch.ResponseServiceDependencies{
			NameCleaner:              *deps.NameCleanerModule.NameCleaner,
			SeriesNameMatcherService: *deps.SeriesNameMatcherModule.SeriesNameMatcherService,
		},
	)

	return siamintershopProductSearchModule{
		ProductSearchScraper: scraper,
		ResponseService:      responseService,
	}
}
