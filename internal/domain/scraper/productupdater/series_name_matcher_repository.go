package productupdater

type SeriesNameMatcherRepository interface {
	FindByName(string) (*SeriesNameMatcher, error)
	Persist(SeriesNameMatcher) error
}
