package productupdater

type SeriesFinderService struct {
	seriesNameMatcherRepository SeriesNameMatcherRepository
	seriesRepository            SeriesRepository
}

func (s SeriesFinderService) FindByNameMatcher(snm SeriesNameMatcher) (*Series, error) {
	existingNameMatcher, err := s.seriesNameMatcherRepository.FindByName(snm.Name())
	if err != nil {
		return nil, err
	}
	if existingNameMatcher == nil {
		return nil, nil
	}

	return s.seriesRepository.FindById(existingNameMatcher.SeriesID)
}
