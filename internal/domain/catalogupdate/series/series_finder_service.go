package series

import "github.com/google/uuid"

type SeriesFinderService struct {
	seriesRepository            seriesRepository
	seriesNameMatcherRepository seriesNameMatcherRepository
}

type NewSeriesFinderServiceDependencies struct {
	SeriesRepository            seriesRepository
	SeriesNameMatcherRepository seriesNameMatcherRepository
}

func NewSeriesFinderService(deps NewSeriesFinderServiceDependencies) *SeriesFinderService {
	return &SeriesFinderService{
		seriesRepository:            deps.SeriesRepository,
		seriesNameMatcherRepository: deps.SeriesNameMatcherRepository,
	}
}

func (sfs SeriesFinderService) MatchOrCreateSeriesByName(name string, publisherID uuid.UUID) (*Series, error) {
	snm, err := sfs.seriesNameMatcherRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	if snm == nil {
		return sfs.createSeriesByName(name, publisherID)
	}

	return sfs.seriesRepository.FindById(snm.seriesID)
}

func (sfs SeriesFinderService) createSeriesByName(name string, publisherID uuid.UUID) (*Series, error) {
	series, err := newSeries(newSeriesArgs{publisherID, name})
	if err != nil {
		return nil, err
	}

	err = sfs.seriesRepository.Persist(*series)
	if err != nil {
		return nil, err
	}

	return series, nil
}
