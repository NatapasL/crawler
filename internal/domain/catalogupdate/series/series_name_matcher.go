package series

import "github.com/google/uuid"

type SeriesNameMatcher struct {
	id       uuid.UUID
	name     string
	seriesID uuid.UUID
}

type newSeriesNameMatcherArgs struct {
	name     string
	seriesID uuid.UUID
}

type ExistingSeriesNameMatcherArgs struct {
	ID       uuid.UUID
	Name     string
	SeriesID uuid.UUID
}

func newSeriesNameMatcher(args newSeriesNameMatcherArgs) (*SeriesNameMatcher, error) {
	snm := SeriesNameMatcher{id: uuid.New(), name: args.name}
	err := snm.validate()
	if err != nil {
		return nil, err
	}
	return &snm, nil
}

func ExistingSeriesNameMatcher(args ExistingSeriesNameMatcherArgs) (*SeriesNameMatcher, error) {
	snm := SeriesNameMatcher{
		id:       args.ID,
		name:     args.Name,
		seriesID: args.SeriesID,
	}
	err := snm.validate()
	if err != nil {
		return nil, err
	}

	return &snm, nil
}

func (SeriesNameMatcher) validate() error {
	return nil
}

func (snm SeriesNameMatcher) Name() string {
	return snm.name
}

func (snm SeriesNameMatcher) ID() uuid.UUID {
	return snm.id
}
