package productupdater

import "github.com/google/uuid"

type SeriesRepository interface {
	FindById(uuid.UUID) (*Series, error)
	Persist(Series) (uuid.UUID, error)
}
