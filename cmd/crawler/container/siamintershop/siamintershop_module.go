package siamintershop

import (
	"manga-crawler/cmd/crawler/container/seriesnamematcher"
	"manga-crawler/internal/domain/scraper/siamintershop"
)

type SiamintershopModule struct {
	Scraper *siamintershop.Scraper
}

type SiamintershopModuleDependencies struct {
	SeriesNameMatcherModule seriesnamematcher.SeriesNameMatcherModule
}

func InitializeSiamintershopModule(deps SiamintershopModuleDependencies) SiamintershopModule {
	apiModule := initializeSiamintershopApiModule(siamintershopApiModuleDependency{})
	nameCleanerModule := initializeSiamintershopNameCleanerModule(
		siamintershopNameCleanerDependencies{Patterns: siamintershop.NameCleanerPattern},
	)
	productSearchModule := initializeSiamintershopProductSearchModule(siamintershopProductSearchModuleDependencies{
		ApiModule:               apiModule,
		SeriesNameMatcherModule: deps.SeriesNameMatcherModule,
	})
	productDetailModule := initializeSiamintershopProductDetailModule(siamintershopProductDetailDependencies{
		ApiModule:               apiModule,
		SeriesNameMatcherModule: deps.SeriesNameMatcherModule,
	})
	categoryListModule := initializeSiamintershopCategoryListModule(SiamintershopCategoryListModuleDependencies{
		ApiModule: apiModule,
	})

	siamintershopScraper := siamintershop.NewScraper(siamintershop.ScraperDependencies{
		ProductSearchScraper: *productSearchModule.ProductSearchScraper,
		ProductDetailScraper: *productDetailModule.ProductDetailScraper,
		CategoryListScraper:  *categoryListModule.CategoryListScraper,
		NameCleaner:          nameCleanerModule.NameCleaner,
	})
	return SiamintershopModule{
		Scraper: siamintershopScraper,
	}
}
