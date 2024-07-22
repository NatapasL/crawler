package categorylist

type CategoryListScraper struct {
	gateway GetCategoryListGateway
}

type CategoryListScraperDependencies struct {
	Gateway GetCategoryListGateway
}

func NewCategoryListScraper(deps CategoryListScraperDependencies) *CategoryListScraper {
	return &CategoryListScraper{gateway: deps.Gateway}
}

func (scraper CategoryListScraper) Scrape() ([]Category, error) {
	categories, err := scraper.gateway.Request()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
