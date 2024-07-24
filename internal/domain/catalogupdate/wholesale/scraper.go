package wholesale

type Scraper interface {
	Scrape(next <-chan bool, product chan<- Product) error
}
