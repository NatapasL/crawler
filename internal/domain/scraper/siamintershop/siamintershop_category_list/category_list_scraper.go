package siamintershopcategorylist

type CategoryListScraper struct {
	gateway GetCategoryListGateway
}

func NewCategoryListScraper(gateway GetCategoryListGateway) *CategoryListScraper {
	return &CategoryListScraper{gateway}
}

func (scraper CategoryListScraper) Scrape() ([]CategoryResponse, error) {
	res, err := scraper.gateway.Request()
	if err != nil {
		return nil, err
	}

	return NewCategoryResponseList(res), nil
}
