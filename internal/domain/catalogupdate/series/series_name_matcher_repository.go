package series

type seriesNameMatcherRepository interface {
	FindByName(string) (*SeriesNameMatcher, error)
}
