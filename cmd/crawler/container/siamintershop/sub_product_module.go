package siamintershop

import (
	"manga-crawler/internal/domain/scraper/siamintershop/subproduct"
)

type subProductModule struct {
	SubProductScraper *subproduct.SubProductScraper
}

type subProductModuleDependencies struct {
	ApiModule siamintershopApiModule
}

func initializeSubProductModule(deps subProductModuleDependencies) subProductModule {
	scraper := subproduct.NewSubProductScraper(subproduct.SubProductScraperDependencies{Gateway: deps.ApiModule.GetSubProductsApi})

	return subProductModule{
		SubProductScraper: scraper,
	}
}
