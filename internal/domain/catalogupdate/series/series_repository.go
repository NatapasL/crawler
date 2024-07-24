package series

import "github.com/google/uuid"

type seriesRepository interface {
	FindById(uuid.UUID) (*Series, error)
	Persist(Series) error
}
