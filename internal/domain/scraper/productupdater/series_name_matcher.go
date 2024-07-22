package productupdater

import "github.com/google/uuid"

type SeriesNameMatcher struct {
	id   uuid.UUID
	name string

	SeriesID uuid.UUID
}

type NewSeriesNameMatcherArgs struct {
	Name string
}

func NewSeriesNameMatcher(args NewSeriesNameMatcherArgs) SeriesNameMatcher {
	return SeriesNameMatcher{
		id:   uuid.New(),
		name: args.Name,
	}
}

func (snm SeriesNameMatcher) ID() uuid.UUID {
	return snm.id
}

func (snm SeriesNameMatcher) Name() string {
	return snm.name
}

func (snm SeriesNameMatcher) ToSeries(publisherId uuid.UUID) (*Series, error) {
	return NewSeries(NewSeriesArgs{Name: snm.name, PublisherID: publisherId})
}
