package model

import "github.com/google/uuid"

type SeriesNameMatcher struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SeriesID uuid.UUID `gorm:"type:uuid"`
	Name     string
}

func (SeriesNameMatcher) TableName() string {
	return "series_name_matcher"
}
