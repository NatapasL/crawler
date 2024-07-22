package container

import (
	"manga-crawler/internal/domain/scraper/namecleaner"
	"manga-crawler/internal/domain/scraper/siamintershop"
	siamintershopcategorylist "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_category_list"
	siamintershopproductdetail "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_detail"
	siamintershopproductsearch "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_search"
)

type SiamintershopModule struct {
	Scraper *siamintershop.Scraper
}

type SiamintershopModuleDependencies struct {
	ApiModule               SiamintershopApiModule
	SeriesNameMatcherModule SeriesNameMatcherModule
}

func InitializeSiamintershopModule(deps SiamintershopModuleDependencies) SiamintershopModule {
	regexpPattern := namecleaner.NewRegexpPattern(siamintershop.GetNameCleanerPattern())
	nameCleaner := namecleaner.NewNameCleaner(regexpPattern)

	// product search
	productSearchScraper := siamintershopproductsearch.NewProductSearchScraper(deps.ApiModule.ProductSearchApi)
	productSearchResponseService := siamintershopproductsearch.NewResponseService(
		*nameCleaner,
		*deps.SeriesNameMatcherModule.SeriesNameMatcherService,
	)

	// product detail
	productDetailScraper := siamintershopproductdetail.NewProductDetailScraper(deps.ApiModule.GetProductDetailApi)
	productDetailResponseService := siamintershopproductdetail.NewResponseService(
		*nameCleaner,
		*deps.SeriesNameMatcherModule.SeriesNameMatcherService,
	)

	// category list
	categoryListScraper := siamintershopcategorylist.NewCategoryListScraper(deps.ApiModule.GetCategoryListApi)

	siamintershopScraper := siamintershop.NewScraper(
		*productSearchScraper,
		*productSearchResponseService,
		*productDetailScraper,
		*productDetailResponseService,
		*categoryListScraper,
	)

	return SiamintershopModule{
		Scraper: siamintershopScraper,
	}
}
