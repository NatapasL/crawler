package seriesnamematcher

import (
	"manga-crawler/internal/infrastructure/model"
	"time"

	"github.com/google/uuid"
)

type SeriesNameMatcherToSeriesMapper struct {
	seriesNameMatcherToSeriesNameMapper SeriesNameMatcherToSeriesNameMapper
}

type SeriesNameMatcherToSeriesMapperDependencies struct {
	SeriesNameMatcherToSeriesNameMapper SeriesNameMatcherToSeriesNameMapper
}

func NewSeriesNameMatcherToSeriesMapper(deps SeriesNameMatcherToSeriesMapperDependencies) *SeriesNameMatcherToSeriesMapper {
	return &SeriesNameMatcherToSeriesMapper{seriesNameMatcherToSeriesNameMapper: deps.SeriesNameMatcherToSeriesNameMapper}
}

func (mapper SeriesNameMatcherToSeriesMapper) Map(nameMatcher model.SeriesNameMatcher, publisherId uuid.UUID) model.Series {
	name := mapper.seriesNameMatcherToSeriesNameMapper.Map(nameMatcher)

	series := model.Series{
		ID:          uuid.New(),
		PublisherID: publisherId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Names:       []model.SeriesName{name},
	}

	return series
}
