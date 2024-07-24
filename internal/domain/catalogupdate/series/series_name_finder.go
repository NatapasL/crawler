package series

import "github.com/google/uuid"

type SeriesNameFinder interface {
	MatchOrCreateSeriesByName(name string, publisherID uuid.UUID) (*Series, error)
}
