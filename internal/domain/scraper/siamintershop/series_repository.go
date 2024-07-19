package siamintershop

import (
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
)

type SeriesRepository interface {
	FindByName(string) []model.Series
	Persist(series model.Series) (*uuid.UUID, error)
}
