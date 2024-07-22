package container

import siamintershopproductdetail "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_detail"

type siamintershopProductDetailModule struct {
	ProductDetailScraper *siamintershopproductdetail.ProductDetailScraper
	ResponseService      *siamintershopproductdetail.ResponseService
}

type siamintershopProductDetailDependencies struct {
	ApiModule               siamintershopApiModule
	NameCleanerModule       siamintershopNameCleanerModule
	SeriesNameMatcherModule SeriesNameMatcherModule
}

func initializeSiamintershopProductDetailModule(
	deps siamintershopProductDetailDependencies,
) siamintershopProductDetailModule {
	scraper := siamintershopproductdetail.NewProductDetailScraper(
		siamintershopproductdetail.ProductDetailScraperDependencies{Gateway: deps.ApiModule.GetProductDetailApi},
	)
	responseService := siamintershopproductdetail.NewResponseService(
		siamintershopproductdetail.ResponseServiceDependencies{
			NameCleaner:              *deps.NameCleanerModule.NameCleaner,
			SeriesNameMatcherService: *deps.SeriesNameMatcherModule.SeriesNameMatcherService,
		},
	)

	return siamintershopProductDetailModule{
		ProductDetailScraper: scraper,
		ResponseService:      responseService,
	}
}
