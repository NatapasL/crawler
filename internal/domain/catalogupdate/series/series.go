package series

import (
	"manga-crawler/internal/domain/language"

	"github.com/google/uuid"
)

type Series struct {
	id           uuid.UUID
	names        []seriesName
	nameMatchers []SeriesNameMatcher
	publisherID  uuid.UUID
}

type newSeriesArgs struct {
	publisherID uuid.UUID
	name        string
}

func newSeries(args newSeriesArgs) (*Series, error) {
	name, err := newSeriesName(newSeriesNameArgs{
		name: args.name, language: language.Th, isDefault: true,
	})
	if err != nil {
		return nil, err
	}
	series := Series{
		id:          uuid.New(),
		names:       []seriesName{*name},
		publisherID: args.publisherID,
	}

	err = series.addNameMatcher(args.name)
	if err != nil {
		return nil, err
	}

	err = series.validate()
	if err != nil {
		return nil, err
	}

	return &series, nil
}

type ExistingSeriesArgs struct {
	ID           uuid.UUID
	PublisherID  uuid.UUID
	NameMatchers []ExistingSeriesNameMatcherArgs
	Names        []ExistingSeriesNameArgs
}

func ExistingSeries(args ExistingSeriesArgs) (*Series, error) {
	var nameMatchers []SeriesNameMatcher
	for _, nameMatcherArgs := range args.NameMatchers {
		snm, err := ExistingSeriesNameMatcher(ExistingSeriesNameMatcherArgs{
			ID:       nameMatcherArgs.ID,
			Name:     nameMatcherArgs.Name,
			SeriesID: args.ID,
		})
		if err != nil {
			return nil, err
		}
		nameMatchers = append(nameMatchers, *snm)
	}

	var names []seriesName
	for _, nameArgs := range args.Names {
		name, err := existingSeriesName(ExistingSeriesNameArgs{
			ID:        nameArgs.ID,
			Name:      nameArgs.Name,
			Language:  nameArgs.Language,
			IsDefault: nameArgs.IsDefault,
		})
		if err != nil {
			return nil, err
		}
		names = append(names, *name)
	}

	series := Series{
		id:           args.ID,
		names:        names,
		nameMatchers: nameMatchers,
		publisherID:  args.PublisherID,
	}
	err := series.validate()
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

func (s Series) PublisherID() uuid.UUID {
	return s.publisherID
}

func (s Series) Names() []seriesName {
	return s.names
}

func (s Series) NameMatchers() []SeriesNameMatcher {
	return s.nameMatchers
}
