package productupdater

import (
	"errors"

	"github.com/google/uuid"
)

type Series struct {
	id          uuid.UUID
	name        string
	publisherID uuid.UUID
}

type NewSeriesArgs struct {
	Name        string
	PublisherID uuid.UUID
}

func NewSeries(args NewSeriesArgs) (*Series, error) {
	series := Series{id: uuid.New(), name: args.Name, publisherID: args.PublisherID}
	if err := series.validate(); err != nil {
		return nil, err
	}

	return &series, nil
}

func (s Series) ID() uuid.UUID {
	return s.id
}

func (s Series) validate() error {
	err := uuid.Validate(s.id.String())
	if err != nil {
		return errors.New("series id invalid")
	}

	if s.name == "" {
		return errors.New("series name is empty")
	}

	err = uuid.Validate(s.publisherID.String())
	if err != nil {
		return errors.New("series publisher id invalid")
	}

	return nil
}
