package seriesnamematcher

import (
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
)

type SeriesNameMatcherToSeriesNameMapper struct{}

func NewSeriesNameMatcherToSeriesNameMapper() *SeriesNameMatcherToSeriesNameMapper {
	return &SeriesNameMatcherToSeriesNameMapper{}
}

func (SeriesNameMatcherToSeriesNameMapper) Map(snm model.SeriesNameMatcher) model.SeriesName {
	return model.SeriesName{
		ID:          uuid.New(),
		Name:        snm.Name,
		Language:    "TH",
		DefaultName: false,
	}
}
