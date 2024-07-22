package container

import "manga-crawler/internal/domain/scraper/siamintershop/categorylist"

type SiamintershopCategoryListModule struct {
	CategoryListScraper *categorylist.CategoryListScraper
}

type SiamintershopCategoryListModuleDependencies struct {
	ApiModule siamintershopApiModule
}

func initializeSiamintershopCategoryListModule(
	deps SiamintershopCategoryListModuleDependencies,
) SiamintershopCategoryListModule {
	categoryListScraper := categorylist.NewCategoryListScraper(
		categorylist.CategoryListScraperDependencies{
			Gateway: deps.ApiModule.GetCategoryListApi,
		},
	)
	return SiamintershopCategoryListModule{
		CategoryListScraper: categoryListScraper,
	}
}
