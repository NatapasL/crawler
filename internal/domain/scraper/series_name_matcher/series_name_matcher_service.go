package seriesnamematcher

import (
	seriesnamematcherrepository "manga-crawler/internal/domain/scraper/series_name_matcher/series_name_matcher_repository"
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
)

type SeriesNameMatcherService struct {
	seriesNameMatcherRepository     seriesnamematcherrepository.SeriesNameMatcherRepository
	seriesRepository                seriesnamematcherrepository.SeriesRepository
	seriesNameMatcherToSeriesMapper SeriesNameMatcherToSeriesMapper
}

type SeriesNameMatcherServiceDependencies struct {
	SeriesNameMatcherRepository     seriesnamematcherrepository.SeriesNameMatcherRepository
	SeriesRepository                seriesnamematcherrepository.SeriesRepository
	SeriesNameMatcherToSeriesMapper SeriesNameMatcherToSeriesMapper
}

func NewSeriesNameMatcherService(deps SeriesNameMatcherServiceDependencies) *SeriesNameMatcherService {
	return &SeriesNameMatcherService{
		seriesNameMatcherRepository:     deps.SeriesNameMatcherRepository,
		seriesRepository:                deps.SeriesRepository,
		seriesNameMatcherToSeriesMapper: deps.SeriesNameMatcherToSeriesMapper,
	}
}

func (service SeriesNameMatcherService) CreateIfNotExists(nameMatcher model.SeriesNameMatcher, publisherId uuid.UUID) error {
	isExists := service.isExists(nameMatcher)

	if isExists {
		return nil
	}

	service.seriesNameMatcherRepository.Persist(nameMatcher)

	err := service.addSeries(nameMatcher)
	if err != nil {
		return err
	}

	return nil
}

func (service SeriesNameMatcherService) isExists(nameMatcher model.SeriesNameMatcher) bool {
	nameMatchers := service.seriesNameMatcherRepository.FindByName(nameMatcher.Name)

	return len(nameMatchers) > 0
}

func (service SeriesNameMatcherService) addSeries(nameMatcher model.SeriesNameMatcher) error {
	publisherId, _ := uuid.FromBytes([]byte("C941BB52-50BA-4273-81D6-2A56831B5B3A"))
	series := service.seriesNameMatcherToSeriesMapper.Map(nameMatcher, publisherId)

	seriesId, err := service.seriesRepository.Persist(series)
	if err != nil {
		return err
	}

	service.attachSeriesToNameMatcher(*seriesId, nameMatcher.ID)

	return nil
}

func (service SeriesNameMatcherService) attachSeriesToNameMatcher(seriesId uuid.UUID, nameMatcherId uuid.UUID) error {
	nameMatcher, err := service.seriesNameMatcherRepository.FindById(nameMatcherId)
	if err != nil {
		return err
	}

	nameMatcher.SetSeriesID(seriesId)
	service.seriesNameMatcherRepository.Persist(nameMatcher)
	return nil
}
