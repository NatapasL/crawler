package series

import "github.com/google/uuid"

type Series struct {
	id           uuid.UUID
	name         string
	nameMatchers []SeriesNameMatcher
	publisherID  uuid.UUID
}

type newSeriesArgs struct {
	publisherID uuid.UUID
	name        string
}

func newSeries(args newSeriesArgs) (*Series, error) {
	series := Series{
		id:          uuid.New(),
		name:        args.name,
		publisherID: args.publisherID,
	}

	err := series.addNameMatcher(args.name)
	if err != nil {
		return nil, err
	}

	err = series.validate()
	if err != nil {
		return nil, err
	}

	return &series, nil
}

func (s *Series) addNameMatcher(name string) error {
	snm, err := newSeriesNameMatcher(newSeriesNameMatcherArgs{name: name, seriesID: s.id})
	if err != nil {
		return err
	}

	s.nameMatchers = append(s.nameMatchers, *snm)
	return nil
}

func (Series) validate() error {
	return nil
}

func (s Series) ID() uuid.UUID {
	return s.id
}
