package model

import (
	"time"

	"github.com/google/uuid"
)

type SeriesNameMatcher struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SeriesID  uuid.UUID `gorm:"type:uuid"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SeriesNameMatcher) TableName() string {
	return "series_name_matcher"
}

func (snm *SeriesNameMatcher) SetSeriesID(id uuid.UUID) {
	snm.SeriesID = id
}
