package series

import (
	"manga-crawler/internal/domain/language"

	"github.com/google/uuid"
)

type seriesName struct {
	id        uuid.UUID
	name      string
	language  language.Language
	isDefault bool
}

type newSeriesNameArgs struct {
	name      string
	language  language.Language
	isDefault bool
}

func newSeriesName(args newSeriesNameArgs) (*seriesName, error) {
	sn := seriesName{
		id:        uuid.New(),
		name:      args.name,
		language:  args.language,
		isDefault: args.isDefault,
	}
	err := sn.validate()
	if err != nil {
		return nil, err
	}

	return &sn, nil
}

type ExistingSeriesNameArgs struct {
	ID        uuid.UUID
	Name      string
	Language  language.Language
	IsDefault bool
}

func existingSeriesName(args ExistingSeriesNameArgs) (*seriesName, error) {
	sn := seriesName{
		id:        args.ID,
		name:      args.Name,
		language:  args.Language,
		isDefault: args.IsDefault,
	}
	err := sn.validate()
	if err != nil {
		return nil, err
	}

	return &sn, nil
}

func (s seriesName) validate() error {
	return nil
}

func (s seriesName) ID() uuid.UUID {
	return s.id
}

func (s seriesName) Name() string {
	return s.name
}

func (s seriesName) Language() language.Language {
	return s.language
}

func (s seriesName) IsDefault() bool {
	return s.isDefault
}
