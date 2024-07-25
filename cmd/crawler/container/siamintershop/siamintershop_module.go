package siamintershop

import (
	"manga-crawler/internal/domain/catalogupdate/series"
	"manga-crawler/internal/domain/scraper/siamintershop"
)

type SiamintershopModule struct {
	Scraper     *siamintershop.Scraper
	WholesaleID string
}

type SiamintershopModuleDependencies struct {
	SeriesFinder series.SeriesFinderService
}

func InitializeSiamintershopModule(deps SiamintershopModuleDependencies) SiamintershopModule {
	apiModule := initializeSiamintershopApiModule(siamintershopApiModuleDependency{})
	nameCleanerModule := initializeSiamintershopNameCleanerModule(
		siamintershopNameCleanerDependencies{Patterns: siamintershop.NameCleanerPattern},
	)
	productSearchModule := initializeSiamintershopProductSearchModule(siamintershopProductSearchModuleDependencies{
		ApiModule: apiModule,
	})
	productDetailModule := initializeSiamintershopProductDetailModule(siamintershopProductDetailDependencies{
		ApiModule: apiModule,
	})
	categoryListModule := initializeSiamintershopCategoryListModule(SiamintershopCategoryListModuleDependencies{
		ApiModule: apiModule,
	})
	subProductModule := initializeSubProductModule(subProductModuleDependencies{
		ApiModule: apiModule,
	})

	siamintershopScraper := siamintershop.NewScraper(siamintershop.ScraperDependencies{
		ProductSearchScraper: *productSearchModule.ProductSearchScraper,
		ProductDetailScraper: *productDetailModule.ProductDetailScraper,
		CategoryListScraper:  *categoryListModule.CategoryListScraper,
		NameCleaner:          nameCleanerModule.NameCleaner,
		SubProductScraper:    *subProductModule.SubProductScraper,
		SeriesFinder:         deps.SeriesFinder,
	})
	return SiamintershopModule{
		Scraper:     siamintershopScraper,
		WholesaleID: siamintershop.SiamintershopID,
	}
}
