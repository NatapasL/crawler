package model

import (
	"github.com/google/uuid"
)

type SeriesName struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SeriesID    uuid.UUID `gorm:"type:uuid"`
	Name        string
	Language    string
	DefaultName bool
}

func (SeriesName) TableName() string {
	return "series_name"
}
