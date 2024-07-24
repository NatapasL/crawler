package siamintershop

import (
	"manga-crawler/cmd/crawler/container/seriesnamematcher"
	"manga-crawler/internal/domain/scraper/siamintershop/productdetail"
)

type siamintershopProductDetailModule struct {
	ProductDetailScraper *productdetail.ProductDetailScraper
}

type siamintershopProductDetailDependencies struct {
	ApiModule               siamintershopApiModule
	SeriesNameMatcherModule seriesnamematcher.SeriesNameMatcherModule
}

func initializeSiamintershopProductDetailModule(
	deps siamintershopProductDetailDependencies,
) siamintershopProductDetailModule {
	scraper := productdetail.NewProductDetailScraper(
		productdetail.ProductDetailScraperDependencies{Gateway: deps.ApiModule.GetProductDetailApi},
	)

	return siamintershopProductDetailModule{
		ProductDetailScraper: scraper,
	}
}
