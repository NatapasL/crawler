package wholesale

type Scraper interface {
	Scrape() ([]Product, error)
}
