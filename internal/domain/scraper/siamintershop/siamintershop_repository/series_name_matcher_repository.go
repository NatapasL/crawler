package siamintershoprepository

import (
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
)

type SeriesNameMatcherRepository interface {
	FindByName(string) []model.SeriesNameMatcher
	Persist(model.SeriesNameMatcher) uuid.UUID
	FindById(id uuid.UUID) (model.SeriesNameMatcher, error)
}
