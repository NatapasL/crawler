package siamintershop

import (
	"manga-crawler/internal/domain/scraper/siamintershop/productsearch"
)

type siamintershopProductSearchModule struct {
	ProductSearchScraper *productsearch.ProductSearchScraper
}

type siamintershopProductSearchModuleDependencies struct {
	ApiModule siamintershopApiModule
}

func initializeSiamintershopProductSearchModule(
	deps siamintershopProductSearchModuleDependencies,
) siamintershopProductSearchModule {
	scraper := productsearch.NewProductSearchScraper(
		productsearch.ProductSearchScraperDependencies{Gateway: deps.ApiModule.ProductSearchApi},
	)

	return siamintershopProductSearchModule{
		ProductSearchScraper: scraper,
	}
}
