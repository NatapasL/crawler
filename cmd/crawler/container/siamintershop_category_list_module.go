package container

import siamintershopcategorylist "manga-crawler/internal/domain/scraper/siamintershop/siamintershop_category_list"

type SiamintershopCategoryListModule struct {
	CategoryListScraper *siamintershopcategorylist.CategoryListScraper
}

type SiamintershopCategoryListModuleDependencies struct {
	ApiModule siamintershopApiModule
}

func initializeSiamintershopCategoryListModule(
	deps SiamintershopCategoryListModuleDependencies,
) SiamintershopCategoryListModule {
	categoryListScraper := siamintershopcategorylist.NewCategoryListScraper(
		siamintershopcategorylist.CategoryListScraperDependencies{
			Gateway: deps.ApiModule.GetCategoryListApi,
		},
	)
	return SiamintershopCategoryListModule{
		CategoryListScraper: categoryListScraper,
	}
}
