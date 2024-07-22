package container

import "manga-crawler/internal/domain/scraper/siamintershop/productdetail"

type siamintershopProductDetailModule struct {
	ProductDetailScraper *productdetail.ProductDetailScraper
	ResponseService      *productdetail.ResponseService
}

type siamintershopProductDetailDependencies struct {
	ApiModule               siamintershopApiModule
	NameCleanerModule       siamintershopNameCleanerModule
	SeriesNameMatcherModule SeriesNameMatcherModule
}

func initializeSiamintershopProductDetailModule(
	deps siamintershopProductDetailDependencies,
) siamintershopProductDetailModule {
	scraper := productdetail.NewProductDetailScraper(
		productdetail.ProductDetailScraperDependencies{Gateway: deps.ApiModule.GetProductDetailApi},
	)
	responseService := productdetail.NewResponseService(
		productdetail.ResponseServiceDependencies{
			NameCleaner:              *deps.NameCleanerModule.NameCleaner,
			SeriesNameMatcherService: *deps.SeriesNameMatcherModule.SeriesNameMatcherService,
		},
	)

	return siamintershopProductDetailModule{
		ProductDetailScraper: scraper,
		ResponseService:      responseService,
	}
}
