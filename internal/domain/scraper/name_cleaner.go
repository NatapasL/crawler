package scraper

type NameCleaner interface {
	Clean(string) string
}
