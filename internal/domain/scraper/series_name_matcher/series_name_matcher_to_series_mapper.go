package seriesnamematcher

import (
	"manga-crawler/internal/infrastructure/model"
	"time"

	"github.com/google/uuid"
)

type SeriesNameMatcherToSeriesMapper struct {
	seriesNameMatcherToSeriesNameMapper SeriesNameMatcherToSeriesNameMapper
}

func NewSeriesNameMatcherToSeriesMapper(
	seriesNameMatcherToSeriesNameMapper SeriesNameMatcherToSeriesNameMapper,
) *SeriesNameMatcherToSeriesMapper {
	return &SeriesNameMatcherToSeriesMapper{seriesNameMatcherToSeriesNameMapper}
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
