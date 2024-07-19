package model

import (
	"github.com/google/uuid"
)

type Genre struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name string

	Series []SeriesGenre `gorm:"foreignKey:GenreID"`
}

func (Genre) TableName() string {
	return "genre"
}
