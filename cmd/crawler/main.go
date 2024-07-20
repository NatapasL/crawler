package main

import (
	"log"
	"manga-crawler/config"
	"manga-crawler/internal/domain/scraper/namecleaner"
	seriesnamematcher "manga-crawler/internal/domain/scraper/series_name_matcher"
	"manga-crawler/internal/domain/scraper/siamintershop"
	siamintershopcategorylist "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_category_list"
	siamintershopproductdetail "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_detail"
	siamintershopproductsearch "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_product_search"
	siamintershopgapi "manga-crawler/internal/infrastructure/api_gateway/siamintershop_api"
	"manga-crawler/internal/infrastructure/database/postgres"
	postgresrepository "manga-crawler/internal/infrastructure/repository/postgres_repository"
)

func main() {
	conf := config.GetConfig()
	db := postgres.GetConnection(conf.Postgres)
	if db == nil {
		log.Println("Cant connect to db...")
		return
	}

	// repository
	seriesRepository := postgresrepository.NewSeriesRepository(db)
	seriesNameMatcherRepository := postgresrepository.NewSeriesNameMatcherRepository(db)

	// series name matcher
	seriesNameMatcherToSeriesNameMapper := seriesnamematcher.NewSeriesNameMatcherToSeriesNameMapper()
	seriesNameMatcherToSeriesMapper := seriesnamematcher.NewSeriesNameMatcherToSeriesMapper(*seriesNameMatcherToSeriesNameMapper)
	seriesNameMatcherService := seriesnamematcher.NewSeriesNameMatcherService(
		seriesNameMatcherRepository,
		seriesRepository,
		*seriesNameMatcherToSeriesMapper,
	)

	// siamintershop
	apiRequest := siamintershopgapi.ApiRequest{}
	productSearchApi := siamintershopgapi.NewProductSearchApi(apiRequest)
	getProductDetailApi := siamintershopgapi.NewGetProductDetailApi(apiRequest)
	getCategoryListApi := siamintershopgapi.NewGetCategoryListApi(apiRequest)

	regexpPattern := namecleaner.NewRegexpPattern(siamintershop.GetNameCleanerPattern())
	nameCleaner := namecleaner.NewNameCleaner(regexpPattern)

	// siamintershop product search
	siamintershopProductSearchScraper := siamintershopproductsearch.NewProductSearchScraper(productSearchApi)
	siamintershopProductSearchResponseService := siamintershopproductsearch.NewResponseService(*nameCleaner, *seriesNameMatcherService)

	// siamintershop product detail
	siamintershopProductDetailScraper := siamintershopproductdetail.NewProductDetailScraper(getProductDetailApi)
	siamintershopProductDetailResponseService := siamintershopproductdetail.NewResponseService(*nameCleaner, *seriesNameMatcherService)

	// siamintershop category list
	siamintershopCategoryListScraper := siamintershopcategorylist.NewCategoryListScraper(getCategoryListApi)

	siamintershopScraper := siamintershop.NewScraper(
		*siamintershopProductSearchScraper,
		*siamintershopProductSearchResponseService,
		*siamintershopProductDetailScraper,
		*siamintershopProductDetailResponseService,
		*siamintershopCategoryListScraper,
	)

	// run
	siamintershopScraper.ScrapeProductDetail()
}
