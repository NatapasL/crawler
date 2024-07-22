package siamintershopcategorylist

type CategoryListScraper struct {
	gateway GetCategoryListGateway
}

type CategoryListScraperDependencies struct {
	Gateway GetCategoryListGateway
}

func NewCategoryListScraper(deps CategoryListScraperDependencies) *CategoryListScraper {
	return &CategoryListScraper{gateway: deps.Gateway}
}

func (scraper CategoryListScraper) Scrape() ([]CategoryResponse, error) {
	res, err := scraper.gateway.Request()
	if err != nil {
		return nil, err
	}

	return NewCategoryResponseList(res), nil
}
