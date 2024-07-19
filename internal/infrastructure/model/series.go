package model

import (
	"time"

	"github.com/google/uuid"
)

type Series struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	PublisherID uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Names     []SeriesName  `gorm:"foreignKey:SeriesID"`
	Genres    []SeriesGenre `gorm:"foreignKey:SeriesID"`
	Publisher Publisher
}
