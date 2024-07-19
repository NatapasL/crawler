package model

import (
	"github.com/google/uuid"
)

type SeriesGenre struct {
	SeriesID uuid.UUID `gorm:"primaryKey;type:uuid"`
	GenreID  uuid.UUID `gorm:"primaryKey;type:uuid"`

	Series Series
	Genre  Genre
}

func (SeriesGenre) TableName() string {
	return "series_genre"
}
