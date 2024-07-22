package productupdater

import "github.com/google/uuid"

type NameMatcherAttacherService struct {
	seriesRepository            SeriesRepository
	seriesNameMatcherRepository SeriesNameMatcherRepository
}

func (service NameMatcherAttacherService) AttachToNonExistingSeries(snm SeriesNameMatcher, publisherId uuid.UUID) (*Series, error) {
	series, err := snm.ToSeries(publisherId)
	if err != nil {
		return nil, err
	}

	_, err = service.seriesRepository.Persist(*series)
	if err != nil {
		return nil, err
	}

	snm.SeriesID = series.ID()
	err = service.seriesNameMatcherRepository.Persist(snm)
	if err != nil {
		return nil, err
	}

	return series, nil
}
