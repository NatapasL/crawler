package seriesnamematcherrepository

import (
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
)

type SeriesRepository interface {
	Persist(model.Series) (*uuid.UUID, error)
}
